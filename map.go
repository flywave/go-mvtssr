package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import "runtime"

type Map struct {
	m *C.struct__mvtssr_map_t
}

func NewMap(fr *HeadlessFrontend, obser *NativeMapObserver, opts *MapOptions, ropts *ResourceOptions) *Map {
	ret := &Map{m: C.mvtssr_new_map(fr.m, obser.m, opts.m, ropts.m)}
	runtime.SetFinalizer(ret, (*Map).free)
	return ret
}

func (t *Map) free() {
	C.mvtssr_map_free(t.m)
}

func (t *Map) SetStyle(st *Style) {
	C.mvtssr_map_set_style(t.m, st.m)
}

func (t *Map) GetCameraOptions(e *EdgeInsets) *CameraOptions {
	ret := &CameraOptions{m: C.mvtssr_map_camera_options(t.m, e.m)}
	runtime.SetFinalizer(ret, (*CameraOptions).free)
	return ret
}

func (t *Map) JumpTo(c *CameraOptions) {
	C.mvtssr_map_jump_to(t.m, c.m)
}

func (t *Map) CameraForLatLngBounds(bounds *LatLngBounds, e *EdgeInsets, bearing, pitch *float64) *CameraOptions {
	ret := &CameraOptions{m: C.mvtssr_map_camera_for_latlng_bounds(t.m, bounds.m, e.m, (*C.double)(bearing), (*C.double)(pitch))}
	runtime.SetFinalizer(ret, (*CameraOptions).free)
	return ret
}

func (t *Map) LatLngBoundsForCamera(opts *CameraOptions) *LatLngBounds {
	ret := &LatLngBounds{m: C.mvtssr_map_camera_latlng_bounds_for_camera(t.m, opts.m)}
	runtime.SetFinalizer(ret, (*LatLngBounds).free)
	return ret
}

func (t *Map) LatLngBoundsForCameraUNWrapped(opts *CameraOptions) *LatLngBounds {
	ret := &LatLngBounds{m: C.mvtssr_map_camera_latlng_bounds_for_camera_unwrapped(t.m, opts.m)}
	runtime.SetFinalizer(ret, (*LatLngBounds).free)
	return ret
}

func (t *Map) SetBounds(op *BoundOptions) {
	C.mvtssr_map_set_bounds(t.m, op.m)
}

func (t *Map) GetBounds() *BoundOptions {
	ret := &BoundOptions{m: C.mvtssr_map_get_bounds(t.m)}
	runtime.SetFinalizer(ret, (*BoundOptions).free)
	return ret
}

func (t *Map) SetNorthOrientation(ori NorthOrientation) {
	C.mvtssr_map_set_north_orientation(t.m, C.uint(ori))
}

func (t *Map) SetConstrainMode(ori ConstrainMode) {
	C.mvtssr_map_set_constrain_mode(t.m, C.uint(ori))
}

func (t *Map) SetViewportMode(ori ViewportMode) {
	C.mvtssr_map_set_viewport_mode(t.m, C.uint(ori))
}

func (t *Map) SetSize(si *Size) {
	C.mvtssr_map_set_size(t.m, si.m)
}

func (t *Map) GetMapOptions() *MapOptions {
	ret := &MapOptions{m: C.mvtssr_map_get_map_options(t.m)}
	runtime.SetFinalizer(ret, (*MapOptions).free)
	return ret
}

func (t *Map) PixelForLatLng(ll *LatLng) *ScreenCoordinate {
	ret := &ScreenCoordinate{m: C.mvtssr_map_pixel_for_latlng(t.m, ll.m)}
	runtime.SetFinalizer(ret, (*ScreenCoordinate).free)
	return ret
}

func (t *Map) LatLngForPixel(sc *ScreenCoordinate) *LatLng {
	ret := &LatLng{m: C.mvtssr_map_latlng_for_pixel(t.m, sc.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func (t *Map) SetDebug(o MapDebugOptions) {
	C.mvtssr_map_set_debug(t.m, C.uint(o))
}

func (t *Map) GetDebug() MapDebugOptions {
	return MapDebugOptions(C.mvtssr_map_get_debug(t.m))
}

func (t *Map) TriggerRepaint() {
	C.mvtssr_map_trigger_repaint(t.m)
}
