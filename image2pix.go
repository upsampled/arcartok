//Package ascart Package Ascii Art Overkill converts an image
//to ascii art for a given `pallet` of ascii characters.
//
//This is done first by running the image through the golang image transformation library `gift`
//The Filter Options string follows the following format `<gift filter name>=<arg1>.<arg2>;<gift filter name>=<arg1>`
//The Filters are run over the image and then the outputed image is sent to the ascii converter
//
//The ascii converter takes the RGB values from the image and converts them to Luma. The Luma is then
//linearly mapped to the given pallet string (IE: luma=0 maps to the first character in teh pallet string)
//
//The command `ascartok` exposes the main method of the library and libascart is a c library wrapper.
//
//Example usage: ascartok -in ./test/gopher.png -plt=" .:OI M" -flt="resize=50,0"
//                .      O M :I     :  .
//               O  I:IIOOIIOO:O I:IIOI :
//             .  O I  O  MM .OIO:I I OI .
//            .  :I OOM MMMMMMM:MOM.. I  .
//          O  IO .. IOMMMMMM:.:O I .:II M
//         I  II I   OMMMMMMO .:O  IMII  .
//       .  : I.M    OMMMMMMI  OO   IO  :
//      .   OO OOII  OMMMMMMM IMO   .I:
//     .  IO:MMMM II IIMMMMMMMM O    OI:
//       II MMMMMM OI OMMMMMMMMOI    MO ..:
//    . OIIMMMMM   :O IO MMMMMOI      III O
//   :  OIMMMMMM: .OO   OOIIOII       M.:  .
//      I IMMMMMI  IIIO.:IIII          II: I
// : :O:MIIMMMMMMMMMO.   :OO           . O
//:.OI MI IMMMMMMMMM:   OIIOO           .MO :
//  :O::  O MMMMMMMOO:OOIII:O           IIO O
// IOO .   OIMMMM OOOIII:O:O            MMI  :
// OMO..    IOIIIOII:IOOIM O             OI: I
//  :MIO      II    OO: M MII             OOO :
//. I:IOI             IIM  II             M.:    I
//    OOIM             O I:I               OO
//       MM             II                 MMII ::
//    : OMO                                  OMMM  .
//    : OIO                                  M:M O
//        M.                                OIM.:
//     . :II                                .II
//     . :OMM                               M I..
//      : III                     OOOI       M:
//      .O I.M                   IIMMII      :O
//       . .IO                   I MMMO      OO  .
//          IIO                 I:OO IO      OIM O
//        .  ..M                OI IOI       IIO .
//         . IOO                             IIO .
//          O : .                            IIO :
//          . :IO                            II: :
//             : M                           II: :
//            :. .                           III :
//           M OIO                           II: :
//           : MOO                           II. M
//              :M                           OI: :
//              : M                          .O  .
//                .                           M
//              O .                         . :
//                .                         O
//              O O                         II: I
//                O                         .O
//              I .                         M:.OI I
//              : M                        MIIMM: O.
//              O .                       O :I MOI .
//            O : I                      I :  M.O M
//            . OII                    MI I
//            O  O                   MI I.  :
//            : :OI.               MI IO:  :   O
//                .I.           M.O I::   .
//                 IIM      M.OI  I.O
//              .:  :M:OO.MM  IIOO    O
//                    I MOIOM::     .
//                 M OMMM:        O
//                O I:MM      II
//                   IM I .
//                O  .OM
//
package ascart

import (
	"image"
	"image/color"
	"os"

	"github.com/disintegration/gift"
)

//AscArt holds the ascii art
type AscArt struct {
	H   int    //Height
	W   int    //Width
	Art string //Art
}

func (a *AscArt) String() string {
	out := ""
	art := []byte(a.Art)
	for i := 0; i < a.H; i++ {
		for j := 0; j < a.W; j++ {
			out += string(art[i*a.W+j])
		}
		out += "\n"
	}
	return out
}

//Img2asc returns an ascii art object for a file name
func Img2asc(fn string, plte string, flarg string) (*AscArt, error) {
	img, err := loadImage(fn)
	if err != nil {
		return nil, err
	}

	flts := procFilter(flarg)
	src := img
	for _, flt := range flts {
		g := gift.New(flt)
		dst := image.NewNRGBA(g.Bounds(src.Bounds()))
		g.Draw(dst, src)
		src = dst
	}

	art := pix2asc(src, plte)
	return art, nil

}

//loadImage helper function to load the file
func loadImage(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

//val2bin given a uint8 value and the number of characters in out pallet
//will return the index of the pallet the value corresponds with
func val2bin(in uint8, bins int) int {
	max := float32(^uint8(0))
	rng := float32(max) / float32(bins)
	val := float32(in)
	lspace := float32(0)
	bin := 0
	for lspace < max {
		if val >= lspace && val < lspace+rng {
			return bin
		}
		bin++
		lspace += rng
	}
	//should never get here
	return bins - 1
}

//pix2asc given an image it will convert each pixel to a character in the given pallet
func pix2asc(im image.Image, plte string) *AscArt {
	out := &AscArt{}
	b := im.Bounds()
	bplt := []byte(plte)

	out.H = b.Max.Y - b.Min.Y
	out.W = b.Max.X - b.Min.X

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, _ := im.At(x, y).RGBA()
			luma, _, _ := color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))
			bin := val2bin(luma, len(plte))
			out.Art += string(bplt[bin])
		}
	}
	return out
}
