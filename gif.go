package aocutils

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func CreateGIF(images []*image.Paletted, palette []color.Color, filename string) {
	var delays []int
	for range images {
		delays = append(delays, 1)
	}

	lastFrame := images[len(images)-1]
	for i := 0; i < 20; i++ {
		images = append(images, lastFrame)
		delays = append(delays, 100)
	}
	gifFile, err := os.Create(filename + ".gif")
	if err != nil {
		panic(err)
	}
	defer gifFile.Close()

	err = gif.EncodeAll(gifFile, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		panic(err)
	}
}
