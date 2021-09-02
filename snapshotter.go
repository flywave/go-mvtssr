package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"

type MapSnapshotter struct {
	m *C.struct__mvtssr_map_snapshotter_t
}

type MapSnapshotterObserver struct {
	m *C.struct__mvtssr_map_snapshotter_observer_t
}

type MapSnapshotterResult struct {
	m *C.struct__mvtssr_map_snapshotter_result_t
}
