package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import "runtime"

type RunLoop struct {
	m *C.struct__mvtssr_runloop_t
}

func NewRunLoop() *RunLoop {
	ret := &RunLoop{m: C.mvtssr_new_runloop()}
	runtime.SetFinalizer(ret, (*RunLoop).free)
	return ret
}

func (t *RunLoop) Stop() {
	C.mvtssr_runloop_free(t.m)
	t.m = nil
}

func (t *RunLoop) free() {
	if t.m != nil {
		C.mvtssr_runloop_free(t.m)
	}
}
