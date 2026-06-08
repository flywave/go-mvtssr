package mvtssr

import (
	"container/list"
	"sync"
	"time"
)

// TileKey uniquely identifies a rendered tile.
type TileKey struct {
	X uint32
	Y uint32
	Z uint32
}

type cacheEntry struct {
	key   TileKey
	value []byte
	exp   time.Time
}

// TileCache is an LRU cache with optional TTL for rendered tile blobs.
type TileCache struct {
	mu         sync.RWMutex
	entries    map[TileKey]*list.Element
	lru        *list.List
	maxEntries int
	ttl        time.Duration
}

// NewTileCache creates a tile blob cache.
// maxEntries <= 0 means unlimited entries (no LRU eviction).
// ttl <= 0 means no expiration.
func NewTileCache(maxEntries int, ttl time.Duration) *TileCache {
	return &TileCache{
		entries:    make(map[TileKey]*list.Element),
		lru:        list.New(),
		maxEntries: maxEntries,
		ttl:        ttl,
	}
}

// Get returns cached tile data, or nil if missing/expired.
func (c *TileCache) Get(key TileKey) []byte {
	c.mu.RLock()
	elem, ok := c.entries[key]
	if !ok {
		c.mu.RUnlock()
		return nil
	}
	entry := elem.Value.(*cacheEntry)
	if !entry.exp.IsZero() && time.Now().After(entry.exp) {
		c.mu.RUnlock()
		c.Delete(key)
		return nil
	}
	c.mu.RUnlock()
	c.mu.Lock()
	c.lru.MoveToFront(elem)
	c.mu.Unlock()
	return entry.value
}

// Set stores tile data in the cache.
func (c *TileCache) Set(key TileKey, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.maxEntries > 0 && c.lru.Len() >= c.maxEntries {
		back := c.lru.Back()
		if back != nil {
			e := back.Value.(*cacheEntry)
			delete(c.entries, e.key)
			c.lru.Remove(back)
		}
	}

	var exp time.Time
	if c.ttl > 0 {
		exp = time.Now().Add(c.ttl)
	}
	entry := &cacheEntry{key: key, value: data, exp: exp}

	if elem, ok := c.entries[key]; ok {
		elem.Value = entry
		c.lru.MoveToFront(elem)
	} else {
		elem := c.lru.PushFront(entry)
		c.entries[key] = elem
	}
}

// Delete removes a cached entry.
func (c *TileCache) Delete(key TileKey) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, ok := c.entries[key]; ok {
		delete(c.entries, key)
		c.lru.Remove(elem)
	}
}

// Len returns the number of cached tiles.
func (c *TileCache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lru.Len()
}

// Clear empties the cache.
func (c *TileCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries = make(map[TileKey]*list.Element)
	c.lru.Init()
}

// RenderCached wraps a tile render callback with caching.
// Rendered tile data is stored in the cache on each callback invocation.
func RenderCached(cache *TileCache, tiler Tiler, start, end uint32, cb func(TileResult) bool) {
	tiler.Render(start, end, func(t TileResult) bool {
		for _, d := range t.GetAllData() {
			if len(d.Image) > 0 {
				cache.Set(TileKey{X: d.X, Y: d.Y, Z: d.Z}, d.Image)
			}
		}
		return cb(t)
	})
}
