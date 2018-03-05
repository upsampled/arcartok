package main

/*
typedef struct{
	int h;
	int w;
	char *art;
}ascart;
*/
import "C"
import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"unsafe"

	"github.com/liamCDI/ascartok"
)

const (
	libAscArtOK C.int = 0 - iota
	libAscArtError
)

//Image2Ascii transform an image to ascii
//export Image2Ascii
func Image2Ascii(cin *C.char, cplt *C.char, cflt *C.char, out **C.ascart) C.int {
	in := C.GoString(cin)
	plt := C.GoString(cplt)
	flt := C.GoString(cflt)
	gout, err := ascart.Img2asc(in, plt, flt)
	if err != nil {
		log.Println(err)
		return libAscArtError
	}
	*out = (*C.ascart)(C.malloc(C.sizeof_ascart)) //Use C calls to malloc or else GB will asasinate traitors
	tmp := (*C.ascart)(unsafe.Pointer(*out))      // If type checking is confused unsafe pointer can correct it
	tmp.h = C.int(gout.H)
	tmp.w = C.int(gout.W)
	tmp.art = C.CString(gout.Art)
	return libAscArtOK
}

func main() {}
