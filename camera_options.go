package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import "runtime"

type CameraOptions struct {
	m *C.struct__mvtssr_camera_options_t
}

func NewCameraOptions() *CameraOptions {
	ret := &CameraOptions{m: C.mvtssr_new_camera_options()}
	runtime.SetFinalizer(ret, (*CameraOptions).free)
	return ret
}

func (t *CameraOptions) free() {
	C.mvtssr_camera_options_free(t.m)
}

func (t *CameraOptions) SetCenter(point *LatLng) *CameraOptions {
	C.mvtssr_camera_options_set_center(t.m, point.m)
	return t
}

func (t *CameraOptions) SetPadding(e *EdgeInsets) *CameraOptions {
	C.mvtssr_camera_options_set_padding(t.m, e.m)
	return t
}

func (t *CameraOptions) SetAnchor(e *ScreenCoordinate) *CameraOptions {
	C.mvtssr_camera_options_set_anchor(t.m, e.m)
	return t
}

func (t *CameraOptions) SetZoom(z float64) *CameraOptions {
	C.mvtssr_camera_options_set_zoom(t.m, C.double(z))
	return t
}

func (t *CameraOptions) SetBearing(b float64) *CameraOptions {
	C.mvtssr_camera_options_set_bearing(t.m, C.double(b))
	return t
}

func (t *CameraOptions) SetPitch(p float64) *CameraOptions {
	C.mvtssr_camera_options_set_pitch(t.m, C.double(p))
	return t
}
