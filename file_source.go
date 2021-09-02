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
	C.mvtssr_file_source_free(t.m)
}

func NewFileSource() *FileSource {
	ret := &FileSource{m: nil}
	ret.m = C.mvtssr_new_file_source(unsafe.Pointer(ret))
	runtime.SetFinalizer(ret, (*FileSource).free)
	return ret
}
