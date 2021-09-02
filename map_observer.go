package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type NativeMapObserver struct {
	m      *C.struct__mvtssr_map_observer_t
	gobser MapObserver
}

func (t *NativeMapObserver) free() {
	C.mvtssr_map_observer_free(t.m)
}

func NullObserver() *NativeMapObserver {
	ret := &NativeMapObserver{m: C.mvtssr_null_map_observer()}
	runtime.SetFinalizer(ret, (*NativeMapObserver).free)
	return ret
}

type MapObserver interface {
	OnCameraWillChange(mode CameraChangeMode)
	OnCameraIsChanging()
	OnCameraDidChange(mode CameraChangeMode)
	OnWillStartLoadingMap()
	OnDidFinishLoadingMap()
	OnDidFailLoadingMap(err MapLoadError, msg string)
	OnWillStartRenderingFrame()
	OnFinishRenderingFrame(mode RenderMode, needsRepaint, placementChanged bool)
	OnWillStartRenderingMap()
	OnDidFinishRenderingMap(mode RenderMode)
	OnDidFinishLoadingStyle()
	OnStyleImageMissing(image string)
	OnDidBecomeIdle()
}

//export goMapObserverOnCameraWillChange
func goMapObserverOnCameraWillChange(ctx unsafe.Pointer, mode C.uint) {
	((*NativeMapObserver)(ctx)).gobser.OnCameraWillChange((CameraChangeMode)(mode))
}

//export goMapObserverOnCameraIsChanging
func goMapObserverOnCameraIsChanging(ctx unsafe.Pointer) {
	((*NativeMapObserver)(ctx)).gobser.OnCameraIsChanging()
}

//export goMapObserverOnCameraDidChange
func goMapObserverOnCameraDidChange(ctx unsafe.Pointer, mode C.uint) {
	((*NativeMapObserver)(ctx)).gobser.OnCameraDidChange((CameraChangeMode)(mode))
}

//export goMapObserverOnWillStartLoadingMap
func goMapObserverOnWillStartLoadingMap(ctx unsafe.Pointer) {
	((*NativeMapObserver)(ctx)).gobser.OnWillStartLoadingMap()
}

//export goMapObserverOnDidFinishLoadingMap
func goMapObserverOnDidFinishLoadingMap(ctx unsafe.Pointer) {
	((*NativeMapObserver)(ctx)).gobser.OnDidFinishLoadingMap()
}

//export goMapObserverOnDidFailLoadingMap
func goMapObserverOnDidFailLoadingMap(ctx unsafe.Pointer, err C.uint, errmsg *C.char) {
	((*NativeMapObserver)(ctx)).gobser.OnDidFailLoadingMap((MapLoadError)(err), C.GoString(errmsg))
}

//export goMapObserverOnWillStartRenderingFrame
func goMapObserverOnWillStartRenderingFrame(ctx unsafe.Pointer) {
	((*NativeMapObserver)(ctx)).gobser.OnWillStartRenderingFrame()
}

//export goMapObserverOnFinishRenderingFrame
func goMapObserverOnFinishRenderingFrame(ctx unsafe.Pointer, mode C.uint, needsRepaint C.bool, placementChanged C.bool) {
	((*NativeMapObserver)(ctx)).gobser.OnFinishRenderingFrame((RenderMode)(mode), bool(needsRepaint), bool(placementChanged))
}

//export goMapObserverOnWillStartRenderingMap
func goMapObserverOnWillStartRenderingMap(ctx unsafe.Pointer) {
	((*NativeMapObserver)(ctx)).gobser.OnWillStartRenderingMap()
}

//export goMapObserverOnDidFinishRenderingMap
func goMapObserverOnDidFinishRenderingMap(ctx unsafe.Pointer, mode C.uint) {
	((*NativeMapObserver)(ctx)).gobser.OnDidFinishRenderingMap((RenderMode)(mode))
}

//export goMapObserverOnDidFinishLoadingStyle
func goMapObserverOnDidFinishLoadingStyle(ctx unsafe.Pointer) {
	((*NativeMapObserver)(ctx)).gobser.OnDidFinishLoadingStyle()
}

//export goMapObserverOnStyleImageMissing
func goMapObserverOnStyleImageMissing(ctx unsafe.Pointer, image *C.char) {
	((*NativeMapObserver)(ctx)).gobser.OnStyleImageMissing(C.GoString(image))
}

//export goMapObserverOnDidBecomeIdle
func goMapObserverOnDidBecomeIdle(ctx unsafe.Pointer) {
	((*NativeMapObserver)(ctx)).gobser.OnDidBecomeIdle()
}
