//go:build linux && egl

package mvtssr

/*
#cgo linux LDFLAGS: -L ./lib/linux -Wl,--start-group -lpthread -lstdc++ -ldl -lm -lEGL -lGLESv2 -ljpeg -lmbgl-core -lmbgl-vendor-csscolorparser -lmbgl-vendor-icu -lmbgl-vendor-nunicode -lmbgl-vendor-parsedate -lmbgl-vendor-sqlite -lmvtssr -lpng -luv -lzlib -Wl,--end-group
*/
import "C"
