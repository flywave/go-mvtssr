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

type ResourceOptions struct {
	m *C.struct__mvtssr_resource_options_t
}

func NewResourceOptions() *ResourceOptions {
	ret := &ResourceOptions{m: C.mvtssr_new_resource_options()}
	runtime.SetFinalizer(ret, (*ResourceOptions).free)
	return ret
}

func (t *ResourceOptions) free() {
	C.mvtssr_resource_options_free(t.m)
}

func (t *ResourceOptions) SetAccessToken(token string) {
	ctoken := C.CString(token)
	defer C.free(unsafe.Pointer(ctoken))
	C.mvtssr_resource_options_set_access_token(t.m, ctoken)
}

func (t *ResourceOptions) SetBaseURL(url string) {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	C.mvtssr_resource_options_set_base_url(t.m, curl)
}

func (t *ResourceOptions) SetAssetPath(path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.mvtssr_resource_options_set_asset_path(t.m, cpath)
}

func (t *ResourceOptions) SetMaximumCacheSize(size uint64) {
	C.mvtssr_resource_options_set_maximum_cache_size(t.m, C.ulong(size))
}
