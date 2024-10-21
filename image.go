package aocutils

import (
	"image"
	"image/color"
)

type Image struct {
	raw  *image.Paletted
	zoom int
	rows int
	cols int
}

func NewImage(rows int, cols int, zoom int) Image {
	return Image{
		raw:  image.NewPaletted(image.Rect(0, 0, cols*zoom, rows*zoom), make([]color.Color, 0, 0)),
		rows: rows,
		cols: cols,
		zoom: zoom,
	}
}

func (img *Image) UsePaletteReds() {
	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		palette = append(palette, color.RGBA{R: 60 + uint8(195*(float64(10-i)/10)), G: 0, B: 0, A: 255})
	}
	palette = append(palette, color.Transparent)
	img.raw.Palette = palette
}

func (img *Image) SetZoomedPixel(i int, j int, value int) {
	for k := i * img.zoom; k < (i+1)*img.zoom; k++ {
		for l := j * img.zoom; l < (j+1)*img.zoom; l++ {
			img.raw.Set(k, l, img.raw.Palette[value])
		}
	}
}
