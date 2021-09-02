package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"

type CameraOptions struct {
	m *C.struct__mvtssr_camera_options_t
}
