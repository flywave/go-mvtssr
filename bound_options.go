package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"

type BoundOptions struct {
	m *C.struct__mvtssr_bound_options_t
}
