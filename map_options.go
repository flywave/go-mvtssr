package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"

type MapOptions struct {
	m *C.struct__mvtssr_map_options_t
}
