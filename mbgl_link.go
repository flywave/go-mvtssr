package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
#cgo linux LDFLAGS: -L ./lib/linux -Wl,--start-group -lpthread -lstdc++ -ldl -lm -lGL -lGLU -lX11 -ljpeg -lmbgl-core -lmbgl-vendor-csscolorparser -lmbgl-vendor-icu -lmbgl-vendor-nunicode -lmbgl-vendor-parsedate -lmbgl-vendor-sqlite -lmvtssr -lpng -luv -lzlib -Wl,--end-group
#cgo darwin LDFLAGS: -L ./lib/darwin -framework CoreFoundation -framework CoreGraphics -framework ImageIO -framework OpenGL -framework CoreText -framework Foundation -ljpeg -lmbgl-core -lmbgl-vendor-csscolorparser -lmbgl-vendor-icu -lmbgl-vendor-nunicode -lmbgl-vendor-parsedate -lmbgl-vendor-sqlite -lmvtssr -lpng -luv -lzlib
#cgo darwin,arm LDFLAGS: -L ./lib/darwin -framework CoreFoundation -framework CoreGraphics -framework ImageIO -framework OpenGL -framework CoreText -framework Foundation -ljpeg -lmbgl-core -lmbgl-vendor-csscolorparser -lmbgl-vendor-icu -lmbgl-vendor-nunicode -lmbgl-vendor-parsedate -lmbgl-vendor-sqlite -lmvtssr -lpng -luv -lzlib
*/
import "C"
