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

type FileSourceFactory struct {
	m *C.struct__mvtssr_file_source_factory_t
}

func NewFileSourceFactory(tp FileType) *FileSourceFactory {
	ret := &FileSourceFactory{m: nil}
	ret.m = C.mvtssr_new_file_source_factory(C.uchar(tp), unsafe.Pointer(ret))
	runtime.SetFinalizer(ret, (*FileSourceFactory).free)
	return ret
}

func (t *FileSourceFactory) free() {
	C.mvtssr_file_source_factory_free(t.m)
}

type FileSourceManager struct {
	m *C.struct__mvtssr_file_source_manager_t
}

func (t *FileSourceManager) free() {
	C.mvtssr_file_source_manager_free(t.m)
}

func NewFileSourceManager() *FileSourceManager {
	ret := &FileSourceManager{m: C.mvtssr_get_file_source_manager()}
	runtime.SetFinalizer(ret, (*FileSourceManager).free)
	return ret
}

func (t *FileSourceManager) Register(f *FileSourceFactory) {
	C.mvtssr_file_source_manager_register_file_source_factory(t.m, f.m)
}

func (t *FileSourceManager) UNRegister(f *FileSourceFactory) {
	C.mvtssr_file_source_manager_unregister_file_source_factory(t.m, f.m)
}
