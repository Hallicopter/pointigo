package pointi

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/EdlinOrg/prominentcolor"
	"github.com/fogleman/gg"
	"github.com/lucasb-eyer/go-colorful"
)

func artistify(imagePath string, saveFlag bool) image.Image {
	img := readImage(imagePath)
	h := img.Bounds().Max.Y
	w := img.Bounds().Max.X
	palette := getPalette(img, 40)

	dc := gg.NewContext(w, h)

	for i := 0; i < w; i += 4 {
		for j := 0; j < h; j += 4 {
			r, g, b, _ := img.At(i, j).RGBA()
			closest := color.RGBA{R: uint8(int(r) + posOrNeg()*rand.Intn(50)),
				G: uint8(int(g) + posOrNeg()*rand.Intn(50)),
				B: uint8(int(b) + posOrNeg()*rand.Intn(50)),
				A: 0xff}

			dc = paintDot(dc, float64(i), float64(j), float64(h/100), getClosestColor(palette, closest))
		}
	}
	if saveFlag {
		dc.SaveJPG("output.jpeg", 80)
	} else {
		return dc.Image()
	}
	return dc.Image()
}

func getPalette(img image.Image, k int) []color.RGBA {
	var palette []color.RGBA

	width := img.Bounds().Max.Y
	out, _ := prominentcolor.KmeansWithAll(k, img, prominentcolor.ArgumentDefault, uint(width)/10, prominentcolor.GetDefaultMasks())

	for _, rgb := range out {

		palette = append(palette, color.RGBA{R: uint8(rgb.Color.R),
			G: uint8(rgb.Color.G),
			B: uint8(rgb.Color.B),
			A: 0xff})
	}

	paletteImg := image.NewRGBA(image.Rect(0, 0, 100*k, 100))
	for i := 0; i < k; i++ {
		for j := 0; j < 100; j++ {
			for l := 0; l < 100; l++ {
				paletteImg.Set(j+100*i, l, palette[i])
			}
		}
	}
	f, err := os.Create("palette.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, paletteImg)

	return palette
}

func paintDot(dc *gg.Context, x float64, y float64, r float64, shade color.RGBA) *gg.Context {
	dc.SetRGBA255(int(shade.R), int(shade.G), int(shade.B), int(shade.A))
	rand := float64(rand.Intn(360))
	dc.RotateAbout(gg.Radians(rand), x, y)
	dc.DrawEllipse(x, y, r, r/0.3)
	dc.RotateAbout(gg.Radians(-rand), x, y)
	dc.Fill()

	return dc
}

func getClosestColor(palette []color.RGBA, shade color.RGBA) color.RGBA {
	var closest color.RGBA
	var minDst float64 = 100

	c1 := colorful.Color{float64(shade.R) / 255.0, float64(shade.G) / 255.0, float64(shade.B) / 255.0}

	for i, clr := range palette {
		dst := c1.DistanceLab(colorful.Color{float64(clr.R) / 255.0,
			float64(clr.G) / 255.0, float64(clr.B) / 255.0})
		// fmt.Println(minDst, dst)
		if dst < minDst {
			closest = palette[i]
			minDst = dst
		}
	}
	return closest
}

func posOrNeg() int {
	n := rand.Intn(1)

	if n == 0 {
		return -1
	}

	return 0
}
