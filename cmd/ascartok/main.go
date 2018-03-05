package main

import (
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/liamCDI/ascartok"
)

func main() {
	in := flag.String("in", "", "file to convert to ascii read")
	flt := flag.String("flt", "resize=100,0", "Go gift filters to be used in the format '<gift filter name>=<arg1>,<arg2>;<gift filter name>...")
	plt := flag.String("plt", " .:OI#M", "pallete string to be used instead of pixels, order for lowest to highest luma")
	flag.Parse()

	out, err := ascart.Img2asc(*in, *plt, *flt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
