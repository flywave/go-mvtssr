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

type FileSource struct {
	m *C.struct__mvtssr_file_source_t
}

func (t *FileSource) free() {
	if t.m != nil {
		C.mvtssr_file_source_free(t.m)
	}
}

//export goFileSourceLoadAsync
func goFileSourceLoadAsync(ctx unsafe.Pointer, req *C.struct__mvtssr_file_source_request_t, res *C.struct__mvtssr_resource_t) {
	freq := &FileSourceRequest{m: req}
	runtime.SetFinalizer(freq, (*FileSourceRequest).free)
	fres := &Resource{m: res}
	runtime.SetFinalizer(fres, (*Resource).free)
	(*(*FileLoader)(ctx)).LoadAsync(freq, fres)
}

//export goFileSourcePause
func goFileSourcePause(ctx unsafe.Pointer) {
	(*(*FileLoader)(ctx)).Pause()
}

//export goFileSourceResume
func goFileSourceResume(ctx unsafe.Pointer) {
	(*(*FileLoader)(ctx)).Resume()
}

type FileSourceResponse struct {
	m *C.struct__mvtssr_file_source_response_t
}

func (t *FileSourceResponse) free() {
	C.mvtssr_file_source_response_free(t.m)
}

func NewFileSourceResponse(data []byte) *FileSourceResponse {
	ret := &FileSourceResponse{m: C.mvtssr_new_file_source_response((*C.char)(unsafe.Pointer(&data[0])), C.size_t(len(data)))}
	runtime.SetFinalizer(ret, (*FileSourceResponse).free)
	return ret
}

func NewErrorFileSourceResponse(code ReasonError, msg string) *FileSourceResponse {
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))
	ret := &FileSourceResponse{m: C.mvtssr_new_file_source_error_response(C.uchar(code), cmsg)}
	runtime.SetFinalizer(ret, (*FileSourceResponse).free)
	return ret
}

type FileSourceRequest struct {
	m *C.struct__mvtssr_file_source_request_t
}

func (t *FileSourceRequest) free() {
	C.mvtssr_file_source_request_free(t.m)
}

func (t *FileSourceRequest) SetResponse(resp *FileSourceResponse) {
	C.mvtssr_file_source_request_set_response(t.m, resp.m)
}
