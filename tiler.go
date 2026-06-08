package mvtssr

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"image/png"
	"math"
	"runtime"
	"sync"
)

// TileData represents a single rendered tile.
type TileData struct {
	Image []byte
	X     uint32
	Y     uint32
	Z     uint32
}

func (d *TileData) IsEmpty() bool {
	return len(d.Image) == 0
}

func (d *TileData) Size() int {
	return len(d.Image)
}

// Format returns "png", "jpeg", or "" unknown.
func (d *TileData) Format() string {
	if len(d.Image) < 4 {
		return ""
	}
	switch {
	case d.Image[0] == 0x89 && d.Image[1] == 'P' && d.Image[2] == 'N' && d.Image[3] == 'G':
		return "png"
	case d.Image[0] == 0xFF && d.Image[1] == 0xD8:
		return "jpeg"
	default:
		return ""
	}
}

// TileResult is the interface for accessing rendered tile data.
type TileResult interface {
	Count() int
	GetAllData() []*TileData
	GetData(i int) (*TileData, error)
}

// RasterTile holds one or more rendered tile data entries.
type RasterTile struct {
	data []*TileData
}

func (t *RasterTile) Count() int {
	return len(t.data)
}

func (t *RasterTile) GetAllData() []*TileData {
	return t.data
}

func (t *RasterTile) GetData(i int) (*TileData, error) {
	if i < 0 || i >= len(t.data) {
		return nil, fmt.Errorf("index %d out of range", i)
	}
	return t.data[i], nil
}

// GridType represents the coordinate system of a tiling scheme.
type GridType int

const (
	GlobalGeodetic GridType = iota
	GlobalMercator
	LocalGeodetic
	LocalMercator
)

// TilerGrid defines a tiling coordinate system.
type TilerGrid struct {
	gridType  GridType
	tileSize  uint32
	rootTiles uint16
	bbox      [4]float64
}

func NewGlobalGeodeticGrid(tileSize uint32, rootTiles uint16) *TilerGrid {
	return &TilerGrid{
		gridType:  GlobalGeodetic,
		tileSize:  tileSize,
		rootTiles: rootTiles,
		bbox:      [4]float64{-180, -90, 180, 90},
	}
}

func NewGlobalMercatorGrid(tileSize uint32, rootTiles uint16) *TilerGrid {
	return &TilerGrid{
		gridType:  GlobalMercator,
		tileSize:  tileSize,
		rootTiles: rootTiles,
		bbox:      [4]float64{-180, -85.05112877980659, 180, 85.05112877980659},
	}
}

func NewGlobalGeodeticLocalGrid(tileSize uint32, box [4]float64, rootTiles uint16) *TilerGrid {
	return &TilerGrid{
		gridType:  LocalGeodetic,
		tileSize:  tileSize,
		rootTiles: rootTiles,
		bbox:      box,
	}
}

func NewGlobalMercatorLocalGrid(tileSize uint32, box [4]float64, rootTiles uint16) *TilerGrid {
	return &TilerGrid{
		gridType:  LocalMercator,
		tileSize:  tileSize,
		rootTiles: rootTiles,
		bbox:      box,
	}
}

// TileBounds returns (minLon, minLat, maxLon, maxLat) for the given tile.
func (g *TilerGrid) TileBounds(x, y, z uint32) (minLon, minLat, maxLon, maxLat float64) {
	n := float64(uint64(1<<z) * uint64(g.rootTiles))
	switch g.gridType {
	case GlobalMercator, LocalMercator:
		minLon = float64(x)/n*360.0 - 180.0
		maxLon = float64(x+1)/n*360.0 - 180.0
		minLat = mercatorInv(float64(y+1) / n)
		maxLat = mercatorInv(float64(y) / n)
	case GlobalGeodetic, LocalGeodetic:
		minLon = float64(x)/n*360.0 - 180.0
		maxLon = float64(x+1)/n*360.0 - 180.0
		minLat = 90.0 - float64(y+1)/n*180.0
		maxLat = 90.0 - float64(y)/n*180.0
	}
	return
}

// Extent returns the tile index range (minX, minY, maxX, maxY) at the given zoom.
func (g *TilerGrid) Extent(zoom uint32) (minX, minY, maxX, maxY uint32) {
	n := uint64(1<<zoom) * uint64(g.rootTiles)
	maxX = uint32(n - 1)
	maxY = uint32(n - 1)
	return
}

// TileCount returns the total number of tiles in the zoom range [start, end] inclusive.
func (g *TilerGrid) TileCount(start, end uint32) uint64 {
	var count uint64
	for z := start; z <= end; z++ {
		n := uint64(1<<z) * uint64(g.rootTiles)
		count += n * n
	}
	return count
}

// TileSize returns the pixel dimensions of each tile.
func (g *TilerGrid) TileSize() uint32 {
	return g.tileSize
}

func mercatorInv(yFraction float64) float64 {
	return rad2deg(math.Atan(math.Sinh(math.Pi - 2*math.Pi*yFraction)))
}

func rad2deg(rad float64) float64 {
	return rad * 180.0 / math.Pi
}

// Tiler is the interface for rendering tiles over a zoom range.
type Tiler interface {
	Render(start, end uint32, cb func(TileResult) bool)
	Free()
}

// RasterTilerOption configures a RasterTiler.
type RasterTilerOption func(*RasterTiler)

// WithOutputFormat sets the image encoding format ("png" or "jpeg").
func WithOutputFormat(format string) RasterTilerOption {
	return func(t *RasterTiler) {
		t.format = format
	}
}

// WithJPEGQuality sets the JPEG encoding quality (1-100).
func WithJPEGQuality(quality int) RasterTilerOption {
	return func(t *RasterTiler) {
		t.jpegQuality = quality
	}
}

// WithPixelRatio sets the pixel ratio for rendering (default 1.0).
func WithPixelRatio(ratio float32) RasterTilerOption {
	return func(t *RasterTiler) {
		t.pixelRatio = ratio
	}
}

// RasterTiler renders raster tiles using Mapbox GL Native via MapSnapshotter.
type RasterTiler struct {
	grid         *TilerGrid
	tileSize     uint32
	pixelRatio   float32
	format       string
	jpegQuality  int
	styleURL     string
	resourceOpts *ResourceOptions
	observer     *NativeMapSnapshotterObserver
	runloop      *RunLoop
	snap         *MapSnapshotter
	mux          sync.Mutex
}

// NewRasterTiler creates a RasterTiler with the given style URL and grid.
// resourceOpts may be nil (defaults to NewResourceOptions()).
func NewRasterTiler(styleURL string, grid *TilerGrid, resourceOpts *ResourceOptions, opts ...RasterTilerOption) *RasterTiler {
	t := &RasterTiler{
		grid:         grid,
		tileSize:     grid.TileSize(),
		pixelRatio:   1.0,
		format:       "png",
		jpegQuality:  85,
		styleURL:     styleURL,
		resourceOpts: resourceOpts,
	}
	if t.resourceOpts == nil {
		t.resourceOpts = NewResourceOptions()
	}
	for _, o := range opts {
		o(t)
	}
	runtime.SetFinalizer(t, (*RasterTiler).free)
	return t
}

func (t *RasterTiler) free() {
	t.Free()
}

// Free releases all native resources.
func (t *RasterTiler) Free() {
	t.mux.Lock()
	defer t.mux.Unlock()
	if t.snap != nil {
		t.snap.free()
		t.snap = nil
	}
	if t.observer != nil {
		t.observer.free()
		t.observer = nil
	}
	if t.runloop != nil {
		t.runloop.Stop()
		t.runloop = nil
	}
	if t.resourceOpts != nil {
		t.resourceOpts.free()
		t.resourceOpts = nil
	}
}

// Render renders tiles in the zoom range [start, end] inclusive.
// The callback cb is called for each rendered tile. Return true to continue
// or false to stop early.
func (t *RasterTiler) Render(start, end uint32, cb func(TileResult) bool) {
	t.mux.Lock()
	if t.snap == nil {
		t.runloop = NewRunLoop()
		t.observer = NullMapSnapshotterObserver()
		size := NewSize(t.tileSize, t.tileSize)
		t.snap = NewMapSnapshotter(size, t.pixelRatio, t.resourceOpts, t.observer)
		if t.styleURL != "" {
			t.snap.SetStyleURL(t.styleURL)
		}
	}
	t.mux.Unlock()

	for z := end; z >= start; z-- {
		minX, minY, maxX, maxY := t.grid.Extent(z)
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				tile := t.renderTile(x, y, z)
				if !cb(tile) {
					return
				}
			}
		}
	}
}

func (t *RasterTiler) renderTile(x, y, z uint32) *RasterTile {
	minLon, minLat, maxLon, maxLat := t.grid.TileBounds(x, y, z)

	sw := NewLatLng(minLat, minLon)
	ne := NewLatLng(maxLat, maxLon)
	bounds := HullLatLngBounds(sw, ne)

	t.snap.SetRegion(bounds)
	result := t.snap.Snapshot()
	img := result.GetImage()

	encoded := encodeImage(img, t.format, t.jpegQuality)

	return &RasterTile{
		data: []*TileData{
			{Image: encoded, X: x, Y: y, Z: z},
		},
	}
}

func encodeImage(img *Image, format string, quality int) []byte {
	rgba := img.Image()
	var buf bytes.Buffer
	switch format {
	case "jpeg":
		jpeg.Encode(&buf, rgba, &jpeg.Options{Quality: quality})
	default:
		png.Encode(&buf, rgba)
	}
	return buf.Bytes()
}

// GridIterator iterates over tile coordinates at a given zoom level without rendering.
type GridIterator struct {
	grid *TilerGrid
	zoom uint32
}

// NewGridIterator creates a GridIterator for the given grid and zoom level.
func NewGridIterator(grid *TilerGrid, zoom uint32) *GridIterator {
	return &GridIterator{grid: grid, zoom: zoom}
}

// TraveGridIterator invokes cb for each tile at the iterator's zoom level.
// The callback receives (x, y, minLon, minLat, maxLon, maxLat).
func (g *GridIterator) TraveGridIterator(cb func(x, y uint32, minLon, minLat, maxLon, maxLat float64)) {
	minX, minY, maxX, maxY := g.grid.Extent(g.zoom)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			mlon, mlat, xlon, xlat := g.grid.TileBounds(x, y, g.zoom)
			cb(x, y, mlon, mlat, xlon, xlat)
		}
	}
}

// ContextRender wraps a context into a Tiler render callback.
// If ctx is cancelled, subsequent tiles are skipped.
func ContextRender(ctx context.Context, tiler Tiler, start, end uint32, cb func(TileResult) bool) {
	tiler.Render(start, end, func(t TileResult) bool {
		select {
		case <-ctx.Done():
			return false
		default:
			return cb(t)
		}
	})
}

// ParallelTiler renders tiles across multiple Tiler instances for throughput.
// Each Tiler must have its own independent resources (separate MapSnapshotter/RunLoop).
func ParallelTiler(tilers []Tiler, totalStart, totalEnd uint32, cb func(TileResult) bool) {
	ParallelTilerCtx(context.Background(), tilers, totalStart, totalEnd, cb)
}

// ParallelTilerCtx is like ParallelTiler but with context cancellation.
func ParallelTilerCtx(ctx context.Context, tilers []Tiler, totalStart, totalEnd uint32, cb func(TileResult) bool) {
	if len(tilers) == 0 {
		return
	}
	perTiler := (totalEnd - totalStart) / uint32(len(tilers))
	if perTiler == 0 {
		perTiler = 1
	}

	var wg sync.WaitGroup

	renderOne := func(tiler Tiler, start, end uint32) {
		defer wg.Done()
		ContextRender(ctx, tiler, start, end, cb)
	}

	for i, t := range tilers {
		wg.Add(1)
		start := totalStart + uint32(i)*perTiler
		end := start + perTiler
		if i == len(tilers)-1 {
			end = totalEnd
		}
		if start < totalEnd {
			go renderOne(t, start, end)
		} else {
			wg.Done()
		}
	}
	wg.Wait()
}
