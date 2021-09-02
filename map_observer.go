package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import "runtime"

type MapObserver struct {
	m *C.struct__mvtssr_map_observer_t
}

func (t *MapObserver) free() {
	C.mvtssr_map_observer_free(t.m)
}

func NullObserver() *MapObserver {
	ret := &MapObserver{m: C.mvtssr_null_map_observer()}
	runtime.SetFinalizer(ret, (*MapObserver).free)
	return ret
}
