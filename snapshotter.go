package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

type MapSnapshotter struct {
	m *C.struct__mvtssr_map_snapshotter_t
}

func (t *MapSnapshotter) free() {
	C.mvtssr_map_snapshotter_free(t.m)
}

func NewMapSnapshotter(si *Size, pixelRatio float32, opts *ResourceOptions, obser *NativeMapSnapshotterObserver) *MapSnapshotter {
	ret := &MapSnapshotter{m: C.mvtssr_new_map_snapshotter(si.m, C.float(pixelRatio), opts.m, obser.m)}
	runtime.SetFinalizer(ret, (*MapSnapshotter).free)
	return ret
}

func (t *MapSnapshotter) SetStyleURL(json string) {
	cjson := C.CString(json)
	defer C.free(unsafe.Pointer(cjson))
	C.mvtssr_map_snapshotter_set_style_url(t.m, cjson)
}

func (t *MapSnapshotter) GetStyleURL() string {
	cjson := C.mvtssr_map_snapshotter_get_style_url(t.m)
	defer C.free(unsafe.Pointer(cjson))
	return C.GoString(cjson)
}

func (t *MapSnapshotter) SetStyleJSON(json string) {
	cjson := C.CString(json)
	defer C.free(unsafe.Pointer(cjson))
	C.mvtssr_map_snapshotter_set_style(t.m, cjson)
}

func (t *MapSnapshotter) GetStyleJSON() string {
	cjson := C.mvtssr_map_snapshotter_get_style(t.m)
	defer C.free(unsafe.Pointer(cjson))
	return C.GoString(cjson)
}

func (t *MapSnapshotter) SetSize(si *Size) {
	C.mvtssr_map_snapshotter_set_size(t.m, si.m)
}

func (t *MapSnapshotter) GetSize() *Size {
	ret := &Size{m: C.mvtssr_map_snapshotter_get_size(t.m)}
	runtime.SetFinalizer(ret, (*Size).free)
	return ret
}

func (t *MapSnapshotter) SetCameraOptions(si *CameraOptions) {
	C.mvtssr_map_snapshotter_set_camera_options(t.m, si.m)
}

func (t *MapSnapshotter) GetCameraOptions() *CameraOptions {
	ret := &CameraOptions{m: C.mvtssr_map_snapshotter_get_camera_options(t.m)}
	runtime.SetFinalizer(ret, (*CameraOptions).free)
	return ret
}

func (t *MapSnapshotter) SetRegion(si *LatLngBounds) {
	C.mvtssr_map_snapshotter_set_region(t.m, si.m)
}

func (t *MapSnapshotter) GetRegion() *LatLngBounds {
	ret := &LatLngBounds{m: C.mvtssr_map_snapshotter_get_region(t.m)}
	runtime.SetFinalizer(ret, (*LatLngBounds).free)
	return ret
}

func (t *MapSnapshotter) Cancel() {
	C.mvtssr_map_snapshotter_cancel(t.m)
}

func (t *MapSnapshotter) Snapshot() *MapSnapshotterResult {
	res := NewMapSnapshotterResult()
	C.mvtssr_map_snapshotter_snapshot(t.m, res.m)
	res.wait()
	return res
}

type MapSnapshotterObserver interface {
	OnDidFailLoadingStyle(style string)
	OnDidFinishLoadingStyle()
	OnStyleImageMissing(image string)
}

type NativeMapSnapshotterObserver struct {
	m      *C.struct__mvtssr_map_snapshotter_observer_t
	gobser MapSnapshotterObserver
}

func (t *NativeMapSnapshotterObserver) free() {
	C.mvtssr_map_snapshotter_observer_free(t.m)
}

func NullMapSnapshotterObserver() *NativeMapSnapshotterObserver {
	ret := &NativeMapSnapshotterObserver{m: C.mvtssr_null_map_snapshotter_observer()}
	runtime.SetFinalizer(ret, (*NativeMapSnapshotterObserver).free)
	return ret
}

func NewMapSnapshotterObserver(ob MapSnapshotterObserver) *NativeMapSnapshotterObserver {
	ret := &NativeMapSnapshotterObserver{m: nil, gobser: ob}
	ret.m = C.mvtssr_new_map_snapshotter_observer(unsafe.Pointer(ret))
	runtime.SetFinalizer(ret, (*NativeMapSnapshotterObserver).free)
	return ret
}

type MapSnapshotterResult struct {
	m *C.struct__mvtssr_map_snapshotter_result_t
}

func (t *MapSnapshotterResult) free() {
	C.mvtssr_map_snapshotter_result_free(t.m)
}

func NewMapSnapshotterResult() *MapSnapshotterResult {
	ret := &MapSnapshotterResult{m: nil}
	ret.m = C.mvtssr_new_map_snapshotter_result(unsafe.Pointer(ret))
	runtime.SetFinalizer(ret, (*MapSnapshotterResult).free)
	return ret
}

func (t *MapSnapshotterResult) wait() {
}

func (t *MapSnapshotterResult) GetImage() *Image {
	ret := &Image{m: C.mvtssr_map_snapshotter_result_get_image(t.m)}
	runtime.SetFinalizer(ret, (*Image).free)
	return ret
}

func (t *MapSnapshotterResult) GetError() error {
	err := C.mvtssr_map_snapshotter_result_get_error(t.m)
	if err != nil {
		defer C.free(unsafe.Pointer(err))
		return errors.New(C.GoString(err))
	}
	return nil
}

func (t *MapSnapshotterResult) PixelForLatLng(ll *LatLng) *ScreenCoordinate {
	ret := &ScreenCoordinate{m: C.mvtssr_map_snapshotter_result_pixel_for_latlng(t.m, ll.m)}
	runtime.SetFinalizer(ret, (*ScreenCoordinate).free)
	return ret
}

func (t *MapSnapshotterResult) LatLngForPixel(sc *ScreenCoordinate) *LatLng {
	ret := &LatLng{m: C.mvtssr_map_snapshotter_result_latlng_for_pixel(t.m, sc.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

//export goMapSnapshotterObserverOnDidFailLoadingStyle
func goMapSnapshotterObserverOnDidFailLoadingStyle(ctx unsafe.Pointer, style *C.char) {
	((*NativeMapSnapshotterObserver)(ctx)).gobser.OnDidFailLoadingStyle(C.GoString(style))
}

//export goMapSnapshotterObserverOnDidFinishLoadingStyle
func goMapSnapshotterObserverOnDidFinishLoadingStyle(ctx unsafe.Pointer) {
	((*NativeMapSnapshotterObserver)(ctx)).gobser.OnDidFinishLoadingStyle()
}

//export goMapSnapshotterObserverOnStyleImageMissing
func goMapSnapshotterObserverOnStyleImageMissing(ctx unsafe.Pointer, image *C.char) {
	((*NativeMapSnapshotterObserver)(ctx)).gobser.OnStyleImageMissing(C.GoString(image))
}
