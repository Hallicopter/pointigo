package pointi

import (
	"image"
)

//GenerateRandomArt takes the API key (id) for the Unsplash API using which it downloads
//a random image and outputs either an image or even a Gif.
func GenerateRandomArt(id string, makeGif bool) {
	getRandomImage(id)
	var images []*image.RGBA

	if makeGif {
		for i := 0; i < 10; i++ {
			img := artistify("random.jpeg", false)
			h = img.Bounds().Max.X
			w = img.Bounds().Max.Y
			images = append(images, img.(*image.RGBA))
		}

		generateGif("out.gif", images)
	}

	artistify("random.jpeg", true)
}

//GenerateArtFromImage takes image path as input and
//outputs either an image or even a Gif.
func GenerateArtFromImage(imagePath string, makeGif bool) {

	var images []*image.RGBA

	if makeGif {
		for i := 0; i < 10; i++ {
			img := artistify("random.jpeg", false)
			h = img.Bounds().Max.X
			w = img.Bounds().Max.Y
			images = append(images, img.(*image.RGBA))
		}

		generateGif("out.gif", images)
	}
	artistify(imagePath, true)
}
