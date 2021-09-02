package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
*/
import "C"
import (
	"image"
	"reflect"
	"runtime"
	"unsafe"
)

type Image struct {
	m *C.struct__mvtssr_premultiplied_image_t
}

func NewEmptyImage() *Image {
	ret := &Image{m: C.mvtssr_empty_premultiplied_image()}
	runtime.SetFinalizer(ret, (*Image).free)
	return ret
}

func NewImage(si *Size) *Image {
	ret := &Image{m: C.mvtssr_new_premultiplied_image(si.m)}
	runtime.SetFinalizer(ret, (*Image).free)
	return ret
}

func NewImageWithData(si *Size, data []byte) *Image {
	ret := &Image{m: C.mvtssr_new_premultiplied_image_with_data(si.m, (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))}
	runtime.SetFinalizer(ret, (*Image).free)
	return ret
}

func (t *Image) free() {
	C.mvtssr_premultiplied_image_free(t.m)
}

func (t *Image) Valid() bool {
	return bool(C.mvtssr_premultiplied_image_valid(t.m))
}

func (t *Image) Stride() int64 {
	return int64(C.mvtssr_premultiplied_image_stride(t.m))
}

func (t *Image) Bytes() int64 {
	return int64(C.mvtssr_premultiplied_image_bytes(t.m))
}

func (t *Image) Size() (width, height uint32) {
	var ch C.uint
	var cw C.uint
	C.mvtssr_premultiplied_image_size(t.m, &cw, &ch)
	width = uint32(cw)
	height = uint32(ch)
	return
}

func (t *Image) Data() []uint8 {
	size := t.Bytes()
	data := C.mvtssr_premultiplied_image_data(t.m)

	var bufSlice []uint8
	bufHeader := (*reflect.SliceHeader)((unsafe.Pointer(&bufSlice)))
	bufHeader.Cap = int(size)
	bufHeader.Len = int(size)
	bufHeader.Data = uintptr(unsafe.Pointer(data))

	return bufSlice
}

func (t *Image) Image() *image.RGBA {
	w, h := t.Size()
	image := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	copy(image.Pix, t.Data())
	return image
}
