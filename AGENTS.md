# go-mvtssr

Go bindings for server-side Mapbox Vector Tile rendering via a [mbgln](external/mbgln/) fork of Mapbox GL Native.

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
ANGLE source is fetched automatically by CMake and built via GN + Ninja.
Prerequisites: Git, Python ≥ 3.6, Ninja, GN (downloaded if not in PATH), Windows SDK.

## Architecture

`package mvtssr` — a library, not a binary. Importers get Go types wrapping C pointers (`*C.struct__mvtssr_*_t`).

| Layer | Files | Role |
|-------|-------|------|
| Go wrappers | `*.go` | User-facing API, GC finalizers, `//export` callback trampolines |
| C bridge | `src/mvtssr_c_api.{h,cc}` | C ABI wrapping mbgl C++ objects |
| Native engine | `external/mbgln/` | Mapbox GL Native fork (C++17) |
| Prebuilt libs | `lib/` | `.a` archives per platform, linked via `#cgo LDFLAGS` |

## Key patterns

**CGo bridge**: Every Go type (e.g. `Map`) holds a C pointer `m *C.struct__...` and registers `runtime.SetFinalizer(t, (*Type).free)`. Never manually call `free()` on objects created by library functions.

**Go→C++→Go callbacks**: Go implements interfaces (`MapObserver`, `FileLoader`, `FileSourceCreater`). The address of the Go interface value is passed as `void* ctx` through C into C++. C++ calls back into Go via `//export` functions in the Go files (see `map_observer.go:51-114`, `file_source.go:24-41`, `file_source_manager.go:28-34`, `snapshotter.go:171-189`).

**Async snapshotter**: `MapSnapshotter.Snapshot()` blocks on a Go channel until the C++ engine finishes rendering. Requires a running `RunLoop`.

**RunLoop**: Must be active for async tile/asset loading to work. Create via `NewRunLoop()`.

**Custom file sources**: Implement `FileLoader` (for per-request loading) and/or `FileSourceCreater` (for factory-based resource options). Register through `FileSourceFactory` + `FileSourceManager`.

## Commands

| Command | Purpose |
|---------|---------|
| `mkdir -p build && cmake -S . -B build && cmake --build build` | Build native static lib |
| `go build ./...` | Build Go package |
| `go test ./...` | Run tests (18 unit tests, 52% coverage across all Go types) |
| `go test -cover ./...` | Run tests with coverage report |
| `go vet ./...` | Static analysis |

## Prerequisites (macOS)

```sh
brew install boost
```

## Gotchas

- `mbgl_link*.go` encodes platform-specific linker flags. The `lib/` directory must contain prebuilt `.a` files matching the target OS/arch.
- `go vet` warns about "passing Go type with embedded pointer to C" in `snapshotter.go:131` and similar files. This is intentional — the Go→C++→Go callback pattern requires passing Go interface pointers through C. The code is safe as long as the Go objects outlive the C++ callbacks.
- On Linux, two backends are available: GLX (default, requires X11) and EGL (`FLYWAVE_WITH_EGL=ON`, no X11). Use Go build tag `egl` for the EGL variant: `go build -tags egl ./...`.
- On Windows, the build always uses EGL/ANGLE (Direct3D). ANGLE source is vendored via CMake `FetchContent` into `external/angle/src/` and built with GN + Ninja. Prerequisites: Git, Python, Ninja, GN, Windows SDK.
- `darwin,arm64` (Apple Silicon) uses `lib/darwin_arm/`, `darwin,amd64` (Intel) uses `lib/darwin/`. Build step 1 outputs to the correct directory automatically.
- No external Go dependencies (Go 1.16, pure CGo).
- `BoundOptions.SetCenter` takes `*LatLngBounds` despite the field name — it sets the bounds, not a center point.
- Adding new `//export` functions requires a corresponding C declaration in `src/mvtssr_c_api.h` as `extern` and a C++ bridge class.
