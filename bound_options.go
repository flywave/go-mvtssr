package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import "runtime"

type BoundOptions struct {
	m *C.struct__mvtssr_bound_options_t
}

func NewBoundOptions() *BoundOptions {
	ret := &BoundOptions{m: C.mvtssr_new_bound_options()}
	runtime.SetFinalizer(ret, (*BoundOptions).free)
	return ret
}

func (t *BoundOptions) free() {
	C.mvtssr_bound_options_free(t.m)
}

func (t *BoundOptions) SetCenter(b *LatLngBounds) *BoundOptions {
	C.mvtssr_bound_options_set_bounds(t.m, b.m)
	return t
}

func (t *BoundOptions) SetMinZoom(z float64) *BoundOptions {
	C.mvtssr_bound_options_set_min_zoom(t.m, C.double(z))
	return t
}

func (t *BoundOptions) SetMaxZoom(z float64) *BoundOptions {
	C.mvtssr_bound_options_set_max_zoom(t.m, C.double(z))
	return t
}

func (t *BoundOptions) SetMinPitch(z float64) *BoundOptions {
	C.mvtssr_bound_options_set_min_pitch(t.m, C.double(z))
	return t
}

func (t *BoundOptions) SetMaxPitch(z float64) *BoundOptions {
	C.mvtssr_bound_options_set_max_pitch(t.m, C.double(z))
	return t
}
