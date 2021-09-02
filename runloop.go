package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"

type RunLoop struct {
	m *C.struct__mvtssr_runloop_t
}
