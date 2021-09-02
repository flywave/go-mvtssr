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

type FileLoader interface {
	LoadAsync(req *FileSourceRequest, res *Resource)
	Pause()
	Resume()
}

type FileSourceCreater interface {
	Create(r *ResourceOptions) *C.struct__mvtssr_unique_file_source_t
}

func newFileSource(loader FileLoader) *C.struct__mvtssr_unique_file_source_t {
	return C.mvtssr_new_unique_file_source(unsafe.Pointer(&loader))
}

//export goFileSourceFactoryCreate
func goFileSourceFactoryCreate(ctx unsafe.Pointer, opt *C.struct__mvtssr_resource_options_t) *C.struct__mvtssr_unique_file_source_t {
	ropt := &ResourceOptions{m: opt}
	runtime.SetFinalizer(ropt, (*ResourceOptions).free)
	fsrc := (*(*FileSourceCreater)(ctx)).Create(ropt)
	return fsrc
}

type FileSourceFactory struct {
	m *C.struct__mvtssr_file_source_factory_t
}

func NewFileSourceFactory(tp FileType, creater FileSourceCreater) *FileSourceFactory {
	ret := &FileSourceFactory{m: C.mvtssr_new_file_source_factory(C.uchar(tp), unsafe.Pointer(&creater))}
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

func (t *FileSourceManager) GetFileSource(tp FileType, opt *ResourceOptions) *FileSource {
	ret := &FileSource{m: C.mvtssr_file_source_manager_get_file_source(t.m, C.uchar(tp), opt.m)}
	runtime.SetFinalizer(ret, (*FileSource).free)
	return ret
}

func (t *FileSourceManager) Register(f *FileSourceFactory) {
	C.mvtssr_file_source_manager_register_file_source_factory(t.m, f.m)
}

func (t *FileSourceManager) Unregister(f *FileSourceFactory) {
	C.mvtssr_file_source_manager_unregister_file_source_factory(t.m, f.m)
}
