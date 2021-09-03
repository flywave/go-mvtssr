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

type Resource struct {
	m *C.struct__mvtssr_resource_t
}

func (t *Resource) free() {
	C.mvtssr_resource_free(t.m)
}

func (t *Resource) GetKind() ResourceKind {
	return ResourceKind(C.mvtssr_resource_get_kind(t.m))
}

func (t *Resource) GetUsage() ResourceUsage {
	return ResourceUsage(C.mvtssr_resource_get_usage(t.m))
}

func (t *Resource) SetUsage(u ResourceUsage) {
	C.mvtssr_resource_set_usage(t.m, C.bool(u))
}

func NewStyleResource(url string) *Resource {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	ret := &Resource{m: C.mvtssr_new_resource_style(curl)}
	runtime.SetFinalizer(ret, (*Resource).free)
	return ret
}

func NewSourceResource(url string) *Resource {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	ret := &Resource{m: C.mvtssr_new_resource_source(curl)}
	runtime.SetFinalizer(ret, (*Resource).free)
	return ret
}

func NewTileResource(urltpl string, pixelRatio float32, x, y int32, z int8, isTms bool, m LoadingMethod) *Resource {
	curltpl := C.CString(urltpl)
	defer C.free(unsafe.Pointer(curltpl))
	ret := &Resource{m: C.mvtssr_new_resource_tile(curltpl, C.float(pixelRatio), C.int(x), C.int(y), C.schar(z), C.bool(isTms), C.uchar(m))}
	runtime.SetFinalizer(ret, (*Resource).free)
	return ret
}

func NewGlyphsResource(urltpl string, fontStack string, start, end uint16) *Resource {
	curltpl := C.CString(urltpl)
	defer C.free(unsafe.Pointer(curltpl))
	cfontStack := C.CString(fontStack)
	defer C.free(unsafe.Pointer(cfontStack))
	ret := &Resource{m: C.mvtssr_new_resource_glyphs(curltpl, cfontStack, C.ushort(start), C.ushort(end))}
	runtime.SetFinalizer(ret, (*Resource).free)
	return ret
}

func NewSpriteImageResource(base string, pixelRatio float32) *Resource {
	cbase := C.CString(base)
	defer C.free(unsafe.Pointer(cbase))
	ret := &Resource{m: C.mvtssr_new_resource_sprite_image(cbase, C.float(pixelRatio))}
	runtime.SetFinalizer(ret, (*Resource).free)
	return ret
}

func NewSpriteJSONResource(base string, pixelRatio float32) *Resource {
	cbase := C.CString(base)
	defer C.free(unsafe.Pointer(cbase))
	ret := &Resource{m: C.mvtssr_new_resource_sprite_json(cbase, C.float(pixelRatio))}
	runtime.SetFinalizer(ret, (*Resource).free)
	return ret
}

func NewImageResource(url string) *Resource {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	ret := &Resource{m: C.mvtssr_new_resource_image(curl)}
	runtime.SetFinalizer(ret, (*Resource).free)
	return ret
}
