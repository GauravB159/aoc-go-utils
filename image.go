package aocutils

import "image"

func setZoomedPixel(i int, j int, zoom int, image *image.Paletted, value int) {
	for k := i * zoom; k < (i+1)*zoom; k++ {
		for l := j * zoom; l < (j+1)*zoom; l++ {
			image.Set(k, l, image.Palette[value])
		}
	}
}
