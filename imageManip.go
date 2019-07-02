package pointi

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type imgLinks struct {
	Raw, Full, Regular, Small, Thumb string
}

type imgUrls struct {
	Urls imgLinks
}

var h, w int

func getRandomImage(id string) {

	var list imgUrls
	endpoint := "https://api.unsplash.com/photos/random?client_id=" + id

	fmt.Println("Getting random image from Unsplash....")
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("HTTP request has gophailed smh")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(data), &list)
	}

	response, e := http.Get(list.Urls.Regular)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	file, err := os.Create("random.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Downloaded random image.")
}

func readImage(url string) image.Image {
	infile, err := os.Open(url)
	if err != nil {
		fmt.Println("Couldn't open stolen goods smh")
		panic(err)
	}
	defer infile.Close()

	thumbnail, _, err := image.Decode(infile)
	if err != nil {
		fmt.Println("Big problem with decoding image.")
		panic(err)
	}

	b := thumbnail.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), thumbnail, b.Min, draw.Src)

	return m
}

func generateGif(fileName string, images []*image.RGBA) {

	outGif := &gif.GIF{}
	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	for _, simage := range images {
		palettedImage := image.NewPaletted(image.Rect(0, 0, h, w), palette.Plan9)
		draw.Draw(palettedImage, palettedImage.Rect, simage, image.Rect(0, 0, h, w).Min, draw.Over)
		outGif.Image = append(outGif.Image, palettedImage)
		outGif.Delay = append(outGif.Delay, 1)
	}
	defer f.Close()
	gif.EncodeAll(f, outGif)
}
