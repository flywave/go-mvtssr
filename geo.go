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

func (t *TileID) ScaledTo(z uint8) *TileID {
	ret := &TileID{m: C.mvtssr_canonical_tileid_scaled_to(t.m, C.uchar(z))}
	runtime.SetFinalizer(ret, (*TileID).free)
	return ret
}

func (t *TileID) IsChildOf(a *TileID) bool {
	return bool(C.mvtssr_canonical_tileid_is_child_of(t.m, a.m))
}

func (t *TileID) Less(a *TileID) bool {
	return bool(C.mvtssr_canonical_tileid_is_child_less(t.m, a.m))
}

func (t *TileID) Eq(a *TileID) bool {
	return bool(C.mvtssr_canonical_tileid_eq(t.m, a.m))
}

func (t *TileID) Children() [4]*TileID {
	var ret [4]*TileID
	par := make([]*C.struct__mvtssr_canonical_tileid_t, 4)
	C.mvtssr_canonical_tileid_children(t.m, &par[0])
	for i := range par {
		ret[i] = &TileID{m: par[i]}
		runtime.SetFinalizer(ret[i], (*TileID).free)
	}
	return ret
}

type LatLng struct {
	m *C.struct__mvtssr_latlng_t
}

func NewLatLng(lat, lon float64) *LatLng {
	ret := &LatLng{m: C.mvtssr_new_latlng(C.double(lat), C.double(lon))}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func NewLatLngWithID(id *TileID) *LatLng {
	ret := &LatLng{m: C.new_mvtssr_latlng_with_id(id.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func (t *LatLng) free() {
	C.mvtssr_mvtssr_latlng_free(t.m)
}

func (t *LatLng) GetData() (lat, lon float64) {
	lat = float64(C.mvtssr_latlng_latitude(t.m))
	lon = float64(C.mvtssr_latlng_longitude(t.m))
	return
}

type LatLngBounds struct {
	m *C.struct__mvtssr_latlng_bounds_t
}

func LatLngBoundsWorld() *LatLngBounds {
	ret := &LatLngBounds{m: C.latlng_bounds_world()}
	runtime.SetFinalizer(ret, (*LatLngBounds).free)
	return ret
}

func LatLngBoundsEmpty() *LatLngBounds {
	ret := &LatLngBounds{m: C.latlng_bounds_empty()}
	runtime.SetFinalizer(ret, (*LatLngBounds).free)
	return ret
}

func NewLatLngBounds(l *LatLng) *LatLngBounds {
	ret := &LatLngBounds{m: C.new_latlng_bounds(l.m)}
	runtime.SetFinalizer(ret, (*LatLngBounds).free)
	return ret
}

func HullLatLngBounds(a, b *LatLng) *LatLngBounds {
	ret := &LatLngBounds{m: C.hull_latlng_bounds(a.m, b.m)}
	runtime.SetFinalizer(ret, (*LatLngBounds).free)
	return ret
}

func NewLatLngBoundsWithID(id *TileID) *LatLngBounds {
	ret := &LatLngBounds{m: C.new_latlng_bounds_with_id(id.m)}
	runtime.SetFinalizer(ret, (*LatLngBounds).free)
	return ret
}

func (t *LatLngBounds) free() {
	C.mvtssr_latlng_bounds_free(t.m)
}

func (t *LatLngBounds) Valid() bool {
	return bool(C.latlng_bounds_valid(t.m))
}

func (t *LatLngBounds) South() float64 {
	return float64(C.latlng_bounds_south(t.m))
}

func (t *LatLngBounds) West() float64 {
	return float64(C.latlng_bounds_west(t.m))
}

func (t *LatLngBounds) North() float64 {
	return float64(C.latlng_bounds_north(t.m))
}

func (t *LatLngBounds) East() float64 {
	return float64(C.latlng_bounds_east(t.m))
}

func (t *LatLngBounds) SouthWest() *LatLng {
	ret := &LatLng{m: C.latlng_bounds_southwest(t.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func (t *LatLngBounds) NorthEast() *LatLng {
	ret := &LatLng{m: C.latlng_bounds_northeast(t.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func (t *LatLngBounds) SouthEast() *LatLng {
	ret := &LatLng{m: C.latlng_bounds_southeast(t.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func (t *LatLngBounds) NorthWest() *LatLng {
	ret := &LatLng{m: C.latlng_bounds_northwest(t.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func (t *LatLngBounds) Center() *LatLng {
	ret := &LatLng{m: C.latlng_bounds_center(t.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func (t *LatLngBounds) Constrain(ll *LatLng) *LatLng {
	ret := &LatLng{m: C.latlng_bounds_constrain(t.m, ll.m)}
	runtime.SetFinalizer(ret, (*LatLng).free)
	return ret
}

func (t *LatLngBounds) Extend(ll *LatLng) {
	C.latlng_bounds_extend(t.m, ll.m)
}

func (t *LatLngBounds) ExtendBounds(ll *LatLngBounds) {
	C.latlng_bounds_extend_bounds(t.m, ll.m)
}

func (t *LatLngBounds) Empty() bool {
	return bool(C.latlng_bounds_is_empty(t.m))
}

func (t *LatLngBounds) ContainsTile(id *TileID) bool {
	return bool(C.latlng_bounds_contains_id(t.m, id.m))
}

func (t *LatLngBounds) ContainsPoint(ll *LatLng) bool {
	return bool(C.latlng_bounds_contains_point(t.m, ll.m))
}

func (t *LatLngBounds) ContainsBounds(area *LatLngBounds) bool {
	return bool(C.latlng_bounds_contains_bounds(t.m, area.m))
}

func (t *LatLngBounds) Intersects(area *LatLngBounds) bool {
	return bool(C.latlng_bounds_contains_intersects(t.m, area.m))
}

type EdgeInsets struct {
	m *C.struct__mvtssr_edge_insets_t
}

func NewEdgeInsets(t, l, b, r float64) *EdgeInsets {
	ret := &EdgeInsets{m: C.mvtssr_new_edge_insets(C.double(t), C.double(l), C.double(b), C.double(r))}
	runtime.SetFinalizer(ret, (*EdgeInsets).free)
	return ret
}

func (t *EdgeInsets) free() {
	C.mvtssr_edge_free(t.m)
}

func (t *EdgeInsets) Top() float64 {
	return float64(C.mvtssr_edge_top(t.m))
}

func (t *EdgeInsets) Left() float64 {
	return float64(C.mvtssr_edge_left(t.m))
}

func (t *EdgeInsets) Bottom() float64 {
	return float64(C.mvtssr_edge_bottom(t.m))
}

func (t *EdgeInsets) Right() float64 {
	return float64(C.mvtssr_edge_right(t.m))
}

func (t *EdgeInsets) IsFlush() bool {
	return bool(C.mvtssr_edge_is_flush(t.m))
}

func (t *EdgeInsets) Center(width, height uint16) *ScreenCoordinate {
	ret := &ScreenCoordinate{m: C.mvtssr_edge_get_center(t.m, C.ushort(width), C.ushort(height))}
	runtime.SetFinalizer(ret, (*ScreenCoordinate).free)
	return ret
}

func (t *EdgeInsets) Eq(a *EdgeInsets) bool {
	return bool(C.mvtssr_edge_eq(t.m, a.m))
}

type ScreenCoordinate struct {
	m *C.struct__mvtssr_screen_coordinate_t
}

func NewScreenCoordinate(x, y float64) *ScreenCoordinate {
	ret := &ScreenCoordinate{m: C.mvtssr_new_screen_coordinate(C.double(x), C.double(y))}
	runtime.SetFinalizer(ret, (*ScreenCoordinate).free)
	return ret
}

func (t *ScreenCoordinate) free() {
	C.mvtssr_screen_coordinate_free(t.m)
}

func (t *ScreenCoordinate) Get() (x, y float64) {
	x = float64(C.mvtssr_screen_coordinate_x(t.m))
	y = float64(C.mvtssr_screen_coordinate_y(t.m))
	return
}

type Size struct {
	m *C.struct__mvtssr_size_t
}

func NewSize(width, height uint32) *Size {
	ret := &Size{m: C.mvtssr_new_size(C.uint(width), C.uint(height))}
	runtime.SetFinalizer(ret, (*Size).free)
	return ret
}

func (t *Size) free() {
	C.mvtssr_size_free(t.m)
}

func (t *Size) Area() uint32 {
	return uint32(C.mvtssr_size_area(t.m))
}

func (t *Size) AspectRatio() float32 {
	return float32(C.mvtssr_size_aspect_ratio(t.m))
}

func (t *Size) Empty() bool {
	return bool(C.mvtssr_size_is_empty(t.m))
}
