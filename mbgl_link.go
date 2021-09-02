package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
#cgo linux CXXFLAGS: -I ./lib -std=gnu++14
#cgo darwin CXXFLAGS: -I ./lib -std=gnu++14
#cgo windows CXXFLAGS:  -I ./lib -std=gnu++14
#cgo linux LDFLAGS: -L ./lib/ -Wl,--start-group -lpthread -lstdc++ -ldl -lm -lGL -lGLU -lX11 -ljpeg -lmbgl-core -lmbgl-vendor-csscolorparser -lmbgl-vendor-icu -lmbgl-vendor-nunicode -lmbgl-vendor-parsedate -lmbgl-vendor-sqlite -lmvtssr -lpng -luv -lzlib -Wl,--end-group
#cgo darwin LDFLAGS: -L ./lib/ -framework CoreFoundation -framework CoreGraphics -framework ImageIO -framework OpenGL -framework CoreText -framework Foundation -ljpeg -lmbgl-core -lmbgl-vendor-csscolorparser -lmbgl-vendor-icu -lmbgl-vendor-nunicode -lmbgl-vendor-parsedate -lmbgl-vendor-sqlite -lmvtssr -lpng -luv -lzlib
*/
import "C"
