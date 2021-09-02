package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import "runtime"

type MapOptions struct {
	m *C.struct__mvtssr_map_options_t
}

func (t *MapOptions) free() {
	C.mvtssr_map_options_free(t.m)
}

func NewMapOptions() *MapOptions {
	ret := &MapOptions{m: C.mvtssr_new_map_options()}
	runtime.SetFinalizer(ret, (*MapOptions).free)
	return ret
}

func (t *MapOptions) SetMapMode(mode MapModeType) {
	C.mvtssr_map_options_set_map_mode(t.m, C.uint(mode))
}

func (t *MapOptions) SetConstrainMode(mode ConstrainMode) {
	C.mvtssr_map_options_set_constrain_mode(t.m, C.uint(mode))
}

func (t *MapOptions) SetViewportMode(mode ViewportMode) {
	C.mvtssr_map_options_set_viewport_mode(t.m, C.uint(mode))
}

func (t *MapOptions) SetCrossSourceCollisions(sc bool) {
	C.mvtssr_map_options_set_cross_source_collisions(t.m, C.bool(sc))
}

func (t *MapOptions) SetNorthOrientation(ori NorthOrientation) {
	C.mvtssr_map_options_set_north_orientation(t.m, C.uchar(ori))
}

func (t *MapOptions) SetSize(si *Size) {
	C.mvtssr_map_options_set_size(t.m, si.m)
}

func (t *MapOptions) SetPixelRatio(ratio float32) {
	C.mvtssr_map_options_set_pixel_ratio(t.m, C.float(ratio))
}
