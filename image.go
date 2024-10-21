package aocutils

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image struct {
	raw      *image.Paletted
	zoom     int
	rows     int
	cols     int
	filename string
}

func CreateImage(rows int, cols int, zoom int, filename string) Image {
	return Image{
		raw:      image.NewPaletted(image.Rect(0, 0, cols*zoom, rows*zoom), make([]color.Color, 0, 0)),
		rows:     rows,
		cols:     cols,
		zoom:     zoom,
		filename: filename + ".png",
	}
}

func (img *Image) GetRawImage() *image.Paletted {
	return img.raw
}

func (img *Image) Clone() Image {
	copied_image := CreateImage(img.rows, img.cols, img.zoom, img.filename)
	copied_image.raw = image.NewPaletted(img.raw.Rect, img.raw.Palette)
	copy(copied_image.raw.Pix, img.raw.Pix)
	return copied_image
}

func (img *Image) UsePaletteReds() {
	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		palette = append(palette, color.RGBA{R: 60 + uint8(195*(float64(9-i)/9)), G: 0, B: 0, A: 255})
	}
	palette = append(palette, color.Transparent)
	img.raw.Palette = palette
}

func (img *Image) UsePaletteWideReds() {
	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		palette = append(palette, color.RGBA{R: 0 + uint8(255*(float64(9-i)/9)), G: 0, B: 0, A: 255})
	}
	palette = append(palette, color.Transparent)
	img.raw.Palette = palette
}

func (img *Image) UsePaletteBlues() {
	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		palette = append(palette, color.RGBA{R: 0, G: 0, B: 60 + uint8(195*(float64(9-i)/9)), A: 255})
	}
	palette = append(palette, color.Transparent)
	img.raw.Palette = palette
}

func (img *Image) UsePaletteGreens() {
	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		palette = append(palette, color.RGBA{R: 0, G: 60 + uint8(195*(float64(9-i)/9)), B: 0, A: 255})
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

func (img *Image) WritePNGToFile() {
	outputFile, _ := os.Create(img.filename)
	png.Encode(outputFile, img.raw)
	outputFile.Close()
}
