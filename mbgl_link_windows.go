//go:build windows

package mvtssr

/*
#cgo windows LDFLAGS: -L ./lib/windows -static-libgcc -static-libstdc++ -Wl,-Bstatic -lpthread -Wl,-Bdynamic -lstdc++ -ldl -lm -lEGL -lGLESv2 -ljpeg -lmbgl-core -lmbgl-vendor-csscolorparser -lmbgl-vendor-icu -lmbgl-vendor-nunicode -lmbgl-vendor-parsedate -lmbgl-vendor-sqlite -lmvtssr -lpng -luv -lzlib -lws2_32 -liphlpapi -lpsapi -luserenv -lcrypt32
*/
import "C"
