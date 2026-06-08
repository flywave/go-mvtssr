package mvtssr

/*
#include <stdlib.h>
#include "mvtssr_c_api.h"
#cgo CFLAGS: -I ./lib
#cgo darwin,amd64 LDFLAGS: -L ./lib/darwin -lc++ -framework CoreFoundation -framework CoreGraphics -framework ImageIO -framework OpenGL -framework CoreText -framework Foundation -ljpeg -lmbgl-core -lmbgl-vendor-csscolorparser -lmbgl-vendor-icu -lmbgl-vendor-parsedate -lmvtssr -lpng -luv -lzlib -lsqlite3
#cgo darwin,arm64 LDFLAGS: -L ./lib/darwin_arm -lc++ -framework CoreFoundation -framework CoreGraphics -framework ImageIO -framework OpenGL -framework CoreText -framework Foundation -ljpeg -lmbgl-core -lmbgl-vendor-csscolorparser -lmbgl-vendor-icu -lmbgl-vendor-parsedate -lmvtssr -lpng -luv -lzlib -lsqlite3
*/
import "C"
