package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import "runtime"

type TileID struct {
	m *C.struct__mvtssr_canonical_tileid_t
}

func NewTileID(z uint8, x, y uint32) *TileID {
	ret := &TileID{m: C.new_mvtssr_canonical_tileid(C.uchar(z), C.uint(x), C.uint(y))}
	runtime.SetFinalizer(ret, (*TileID).free)
	return ret
}

func (t *TileID) free() {
	C.mvtssr_canonical_tileid_free(t.m)
}

type LatLng struct {
	m *C.struct__mvtssr_latlng_t
}

type LatLngBounds struct {
	m *C.struct__mvtssr_latlng_bounds_t
}

type EdgeInsets struct {
	m *C.struct__mvtssr_edge_insets_t
}

type ScreenCoordinate struct {
	m *C.struct__mvtssr_screen_coordinate_t
}

type Size struct {
	m *C.struct__mvtssr_size_t
}
