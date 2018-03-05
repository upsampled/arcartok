package ascart

import (
	"image"
	"image/color"
	"log"
	"strconv"
	"strings"

	"github.com/disintegration/gift"
)

func procFilter(flarg string) []gift.Filter {
	out := []gift.Filter{}
	flts := strings.Split(flarg, ";")
	for _, flt := range flts {
		key := strings.Split(flt, "=")
		switch key[0] {
		case "contrast":
			val := strings.Split(key[1], ",")
			c, _ := strconv.ParseFloat(val[0], 32)
			out = append(out, gift.Contrast(float32(c)))
		case "resize":
			val := strings.Split(key[1], ",")
			h, _ := strconv.Atoi(val[0])
			w, _ := strconv.Atoi(val[1])
			out = append(out, gift.Resize(h, w, gift.LanczosResampling))
		case "invert":
			out = append(out, gift.Invert())
		case "sobel":
			out = append(out, gift.Sobel())
		case "crop":
			val := strings.Split(key[1], ",")
			x0, _ := strconv.Atoi(val[0])
			y0, _ := strconv.Atoi(val[1])
			x1, _ := strconv.Atoi(val[2])
			y1, _ := strconv.Atoi(val[3])
			out = append(out, gift.Crop(image.Rect(x0, y0, x1, y1)))
		case "fliphorizontal":
			out = append(out, gift.FlipHorizontal())
		case "flipvertical":
			out = append(out, gift.FlipVertical())
		case "rotate":
			val := strings.Split(key[1], ",")
			c, _ := strconv.ParseFloat(val[0], 32)
			out = append(out, gift.Rotate(float32(c), color.Transparent, gift.CubicInterpolation))
		case "conv":
			val := strings.Split(key[1], ",")
			in := []float32{}
			for i := 0; i < 9; i++ {
				tmp, err := strconv.ParseFloat(val[i], 32)
				if err != nil {
					log.Fatal(err)
				}
				in = append(in, float32(tmp))
			}
			out = append(out, gift.Convolution(in, false, false, false, 0.0))
		}
	}
	return out
}
