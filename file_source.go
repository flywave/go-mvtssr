package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"

type FileSource struct {
	m *C.struct__mvtssr_file_source_t
}
