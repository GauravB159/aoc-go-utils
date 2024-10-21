package aocutils

import (
	"image"
	"image/color"
	"image/png"
	"os"
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

func (img *Image) GetRawImage() *image.Paletted {
	return img.raw
}

func (img *Image) Clone() Image {
	copied_image := NewImage(img.rows, img.cols, img.zoom)
	copied_image.raw = image.NewPaletted(img.raw.Rect, img.raw.Palette)
	copy(copied_image.raw.Pix, img.raw.Pix)
	return copied_image
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

func (img *Image) UsePaletteBlues() {
	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		palette = append(palette, color.RGBA{R: 0, G: 0, B: 60 + uint8(195*(float64(10-i)/10)), A: 255})
	}
	palette = append(palette, color.Transparent)
	img.raw.Palette = palette
}

func (img *Image) UsePaletteGreens() {
	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		palette = append(palette, color.RGBA{R: 0, G: 60 + uint8(195*(float64(10-i)/10)), B: 0, A: 255})
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

func (img *Image) WritePNGToFile(filename string) {
	outputFile, _ := os.Create(filename + ".png")
	png.Encode(outputFile, img.raw)
	outputFile.Close()
}
