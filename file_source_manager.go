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

type FileSourceCreater interface {
	Make(r *ResourceOptions) *FileSource
}

//export goFileSourceFactoryCreate
func goFileSourceFactoryCreate(ctx unsafe.Pointer, opt *C.struct__mvtssr_resource_options_t) *C.struct__mvtssr_file_source_t {
	ropt := &ResourceOptions{m: opt}
	runtime.SetFinalizer(ropt, (*ResourceOptions).free)
	fsrc := ((*FileSourceFactory)(ctx)).creater.Make(ropt)
	m := fsrc.m
	((*FileSourceFactory)(ctx)).fsmap[uintptr(unsafe.Pointer(m))] = fsrc
	return m
}

//export goFileSourceFactoryDestory
func goFileSourceFactoryDestory(ctx unsafe.Pointer, fs unsafe.Pointer) {
	maps := ((*FileSourceFactory)(ctx)).fsmap
	key := uintptr(unsafe.Pointer(fs))
	if fs, ok := maps[key]; ok {
		fs.free()
		delete(maps, key)
	}
}

type FileSourceFactory struct {
	m       *C.struct__mvtssr_file_source_factory_t
	creater FileSourceCreater
	fsmap   map[uintptr]*FileSource
}

func NewFileSourceFactory(tp FileType, creater FileSourceCreater) *FileSourceFactory {
	ret := &FileSourceFactory{m: nil, creater: creater}
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

func (t *FileSourceManager) GetFileSource(tp FileType, opt *ResourceOptions) *FileSource {
	//TODO
	return nil
}

func (t *FileSourceManager) Register(f *FileSourceFactory) {
	C.mvtssr_file_source_manager_register_file_source_factory(t.m, f.m)
}

func (t *FileSourceManager) Unregister(f *FileSourceFactory) {
	C.mvtssr_file_source_manager_unregister_file_source_factory(t.m, f.m)
}
