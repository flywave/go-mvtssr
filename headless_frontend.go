package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"

type HeadlessFrontend struct {
	m *C.struct__mvtssr_headless_frontend_t
}
