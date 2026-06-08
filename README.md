# go-mvtssr

Go bindings for server-side Mapbox Vector Tile rendering via a [mbgln](external/mbgln/) fork of Mapbox GL Native.

Produces raster tiles (PNG/JPEG) from Mapbox GL styles — no display server required.

## Build

Two build systems — **both** must succeed to produce a working binary:

```sh
# 1. Build the C++ static library (requires cmake, C++17 toolchain, OpenGL deps)
mkdir -p build && cmake -S . -B build && cmake --build build

# 2. Build / test the Go package (links against prebuilt .a files in lib/)
go build ./...
go test ./...
```

Output from step 1 lands in `lib/<os>_<arch>/` (e.g. `lib/darwin_arm/`).

For **Linux headless servers** (no X11), build the native lib with EGL instead of GLX:
```sh
cmake -S . -B build -DFLYWAVE_WITH_EGL=ON && cmake --build build
# then build Go with the `egl` tag:
go build -tags egl ./...
```

On **Windows**, the build always uses EGL + ANGLE (Direct3D backend):
```sh
cmake -S . -B build -G "MinGW Makefiles" && cmake --build build
go build ./...
```

## Prerequisites (macOS)

```sh
brew install boost
```

## Usage

### Basic snapshot

```go
package main

import (
	"fmt"
	mvtssr "github.com/flywave/go-mvtssr"
)

func main() {
	loop := mvtssr.NewRunLoop()
	defer loop.Stop()

	ropts := mvtssr.NewResourceOptions()
	ropts.SetAccessToken("mapbox-token")

	obs := mvtssr.NullMapSnapshotterObserver()
	snap := mvtssr.NewMapSnapshotter(mvtssr.NewSize(512, 512), 1.0, ropts, obs)
	snap.SetStyleURL("https://.../style.json")

	result := snap.Snapshot()
	img := result.GetImage()
	if err := result.GetError(); err != nil {
		panic(err)
	}
	fmt.Printf("rendered %dx%d image\n", img.Size())
}
```

### Tile publishing API

The tiler API follows the [flywave-mapnik](https://github.com/flywave/flywave-mapnik) tiler pattern:

```go
grid := mvtssr.NewGlobalMercatorGrid(256, 1)
tiler := mvtssr.NewRasterTiler("https://.../style.json", grid, ropts,
	mvtssr.WithOutputFormat("png"),
)
defer tiler.Free()

tiler.Render(0, 5, func(tile mvtssr.TileResult) bool {
	for _, d := range tile.GetAllData() {
		// d.Image is PNG bytes, d.X/Y/Z are tile coordinates
		saveTile(d.Image, d.Z, d.X, d.Y)
	}
	return true
})
```

### Parallel rendering

```go
tiler1 := mvtssr.NewRasterTiler(styleURL, grid, ropts1)
tiler2 := mvtssr.NewRasterTiler(styleURL, grid, ropts2)
defer tiler1.Free()
defer tiler2.Free()

mvtssr.ParallelTiler([]mvtssr.Tiler{tiler1, tiler2}, 0, 10, cb)
```

### Caching

```go
cache := mvtssr.NewTileCache(1000, 5*time.Minute)
mvtssr.RenderCached(cache, tiler, 0, 10, cb)
```

### Grids

| Function | Description |
|----------|-------------|
| `NewGlobalGeodeticGrid(tileSize, rootTiles)` | EPSG:4326 global grid |
| `NewGlobalMercatorGrid(tileSize, rootTiles)` | EPSG:3857 global grid (Web Mercator) |
| `NewGlobalGeodeticLocalGrid(tileSize, box, rootTiles)` | Bounded EPSG:4326 grid |
| `NewGlobalMercatorLocalGrid(tileSize, box, rootTiles)` | Bounded EPSG:3857 grid |

### Tile iteration (no rendering)

```go
iter := mvtssr.NewGridIterator(grid, 5)
iter.TraveGridIterator(func(x, y uint32, minLon, minLat, maxLon, maxLat float64) {
	fmt.Printf("tile %d/%d: bounds [%f, %f, %f, %f]\n", x, y, minLon, minLat, maxLon, maxLat)
})
```

## API reference

| Go type / function | File | Role |
|---|---|---|
| `Map`, `NewMap` | `map.go` | Full mbgl::Map wrapper |
| `MapSnapshotter`, `Snapshot()` | `snapshotter.go` | Async viewport renderer |
| `RasterTiler`, `Render()` | `tiler.go` | Bulk tile renderer (zoom range) |
| `TilerGrid` | `tiler.go` | Tiling scheme (mercator/geodetic) |
| `TileData`, `TileResult` | `tiler.go` | Rendered tile data + interface |
| `ParallelTiler` | `tiler.go` | Multi-tiler parallel rendering |
| `TileCache`, `RenderCached` | `tilecache.go` | LRU tile blob cache |
| `GridIterator` | `tiler.go` | Tile coordinate iteration |
| `LatLng`, `LatLngBounds` | `geo.go` | Geographic primitives |
| `Image` | `image.go` | Premultiplied RGBA image wrapper |
| `FileSource`, `FileLoader` | `file_source*.go` | Custom file loading |
| `RunLoop` | `runloop.go` | Event loop for async operations |
| `Style` | `style.go` | Map style loader |

## Architecture

```
┌─────────────────┐     ┌──────────────────────┐     ┌──────────────┐
│   User code     │────▶│   Go wrappers (*.go)  │────▶│  C bridge    │
│ (tiler, snap,   │     │ (Map, Snapshotter,    │     │ (src/*.{h,cc})│
│  render, etc.)  │     │  Tiler, Geo types)    │     │              │
└─────────────────┘     └──────────────────────┘     └──────┬───────┘
                                                            │
                                                     ┌──────▼───────┐
                                                     │  mbgln C++   │
                                                     │ (Mapbox GL   │
                                                     │  Native fork)│
                                                     └──────────────┘
```

## Build tags

| Tag | Platform | Backend |
|-----|----------|---------|
| (none) | darwin | OpenGL (CGL) |
| `egl` | linux | EGL (no X11 required) |
| (none) | linux | GLX (requires X11) |
| (none) | windows | EGL + ANGLE (Direct3D) |
