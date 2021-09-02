package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import (
	"runtime"
)

type HeadlessFrontend struct {
	m *C.struct__mvtssr_headless_frontend_t
}

func (t *HeadlessFrontend) free() {
	C.mvtssr_headless_frontend_free(t.m)
}

func NewHeadlessFrontend(si *Size, pixelRatio float32) *HeadlessFrontend {
	ret := &HeadlessFrontend{m: C.mvtssr_new_headless_frontend(si.m, C.float(pixelRatio))}
	runtime.SetFinalizer(ret, (*FileSourceFactory).free)
	return ret
}

func (t *HeadlessFrontend) Reset() {
	C.mvtssr_headless_frontend_reset(t.m)
}

func (t *HeadlessFrontend) PixelForLatLng(ll *LatLng) *ScreenCoordinate {
	ret := &ScreenCoordinate{m: C.mvtssr_headless_frontend_pixel_for_latlng(t.m, ll.m)}
	runtime.SetFinalizer(ret, (*ScreenCoordinate).free)
	return ret
}

func (t *HeadlessFrontend) LatLngForPixel(sc *ScreenCoordinate) *LatLng {
	ret := &LatLng{m: C.mvtssr_headless_frontend_latlng_for_pixel(t.m, sc.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func (t *HeadlessFrontend) SetSize(si *Size) {
	C.mvtssr_headless_frontend_set_size(t.m, si.m)
}

func (t *HeadlessFrontend) GetSize() *Size {
	ret := &Size{m: C.mvtssr_headless_frontend_get_size(t.m)}
	runtime.SetFinalizer(ret, (*Size).free)
	return ret
}

func (t *HeadlessFrontend) Render(m *Map) *Image {
	ret := &Image{m: C.mvtssr_headless_frontend_render(t.m, m.m)}
	runtime.SetFinalizer(ret, (*Image).free)
	return ret
}
