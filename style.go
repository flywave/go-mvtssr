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

type Style struct {
	m *C.struct__mvtssr_style_t
}

func (t *Style) free() {
	C.mvtssr_style_free(t.m)
}

func NewStyle(source *FileSource, pixelRatio float32) *Style {
	ret := &Style{m: C.mvtssr_new_style(source.m, C.float(pixelRatio))}
	runtime.SetFinalizer(ret, (*Style).free)
	return ret
}

func (t *Style) GetJSON() string {
	cjson := C.mvtssr_style_get_json(t.m)
	defer C.free(unsafe.Pointer(cjson))
	return C.GoString(cjson)
}

func (t *Style) GetURL() string {
	curl := C.mvtssr_style_get_url(t.m)
	defer C.free(unsafe.Pointer(curl))
	return C.GoString(curl)
}

func (t *Style) LoadURL(url string) {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	C.mvtssr_style_load_url(t.m, curl)
}

func (t *Style) LoadJSON(json string) {
	cjson := C.CString(json)
	defer C.free(unsafe.Pointer(cjson))
	C.mvtssr_style_load_json(t.m, cjson)
}
