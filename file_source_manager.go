package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"

type FileSourceFactory struct {
	m *C.struct__mvtssr_file_source_factory_t
}

type FileSourceManager struct {
	m *C.struct__mvtssr_file_source_manager_t
}
