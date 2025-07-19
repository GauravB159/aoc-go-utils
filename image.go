package aocutils

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image struct {
	raw      *image.Paletted
	rowzoom  int
	colzoom  int
	rows     int
	cols     int
	filename string
}

func CreateImage(rows int, cols int, rowzoom int, colzoom int, filename string) Image {
	return Image{
		raw:      image.NewPaletted(image.Rect(0, 0, cols*colzoom, rows*rowzoom), make([]color.Color, 0)),
		rows:     rows,
		cols:     cols,
		rowzoom:  rowzoom,
		colzoom:  colzoom,
		filename: filename + ".png",
	}
}

func (img *Image) GetRawImage() *image.Paletted {
	return img.raw
}

func (img *Image) Clone() Image {
	copied_image := CreateImage(img.rows, img.cols, img.rowzoom, img.colzoom, img.filename)
	copied_image.raw = image.NewPaletted(img.raw.Rect, img.raw.Palette)
	copy(copied_image.raw.Pix, img.raw.Pix)
	return copied_image
}

func (img *Image) UsePaletteReds(inverse bool) {
	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		if inverse {
			palette = append(palette, color.RGBA{R: 60 + uint8(195*(float64(i)/9)), G: 0, B: 0, A: 255})
		} else {
			palette = append(palette, color.RGBA{R: 60 + uint8(195*(float64(9-i)/9)), G: 0, B: 0, A: 255})
		}
	}
	palette = append(palette, color.Transparent)
	img.raw.Palette = palette
}

func (img *Image) UsePaletteWideReds(inverse bool) {
	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		if inverse {
			palette = append(palette, color.RGBA{R: 0 + uint8(255*(float64(i)/9)), G: 0, B: 0, A: 255})
		} else {
			palette = append(palette, color.RGBA{R: 0 + uint8(255*(float64(9-i)/9)), G: 0, B: 0, A: 255})
		}
	}
	palette = append(palette, color.Transparent)
	img.raw.Palette = palette
}

func (img *Image) UsePaletteRedToYellow() {
	palette := make([]color.Color, 0, 10)
	palette = append(palette, color.Black)
	palette = append(palette, color.RGBA{R: 255, G: 187, B: 37, A: 255})
	palette = append(palette, color.RGBA{R: 255, G: 170, B: 37, A: 255})
	palette = append(palette, color.RGBA{R: 255, G: 154, B: 37, A: 255})
	palette = append(palette, color.RGBA{R: 255, G: 137, B: 37, A: 255})
	palette = append(palette, color.RGBA{R: 255, G: 121, B: 37, A: 255})
	palette = append(palette, color.RGBA{R: 255, G: 104, B: 37, A: 255})
	palette = append(palette, color.RGBA{R: 255, G: 87, B: 37, A: 255})
	palette = append(palette, color.RGBA{R: 255, G: 71, B: 37, A: 255})
	palette = append(palette, color.RGBA{R: 255, G: 54, B: 37, A: 255})
	palette = append(palette, color.RGBA{R: 255, G: 34, B: 37, A: 255})
	img.raw.Palette = palette
}

func (img *Image) UseFullColors() {
	palette := make([]color.Color, 0, 10)
	palette = append(palette, color.RGBA{R: 0, G: 18, B: 25, A: 255})     // 001219
	palette = append(palette, color.RGBA{R: 0, G: 95, B: 115, A: 255})    // 005F73
	palette = append(palette, color.RGBA{R: 10, G: 147, B: 150, A: 255})  // 0A9396
	palette = append(palette, color.RGBA{R: 148, G: 210, B: 189, A: 255}) // 94D2BD
	palette = append(palette, color.RGBA{R: 233, G: 216, B: 166, A: 255}) // E9D8A6
	palette = append(palette, color.RGBA{R: 238, G: 155, B: 0, A: 255})   // EE9B00
	palette = append(palette, color.RGBA{R: 202, G: 103, B: 2, A: 255})   // CA6702
	palette = append(palette, color.RGBA{R: 187, G: 62, B: 3, A: 255})    // BB3E03
	palette = append(palette, color.RGBA{R: 174, G: 32, B: 18, A: 255})   // AE2012
	palette = append(palette, color.RGBA{R: 155, G: 34, B: 38, A: 255})   // 9B2226
	palette = append(palette, color.RGBA{R: 0, G: 0, B: 0, A: 255})       // 000000
	img.raw.Palette = palette
}

func (img *Image) UsePaletteColors() {
	palette := make([]color.Color, 0, 7)
	palette = append(palette, color.RGBA{R: 38, G: 70, B: 83, A: 255})
	palette = append(palette, color.RGBA{R: 42, G: 157, B: 143, A: 255})
	palette = append(palette, color.RGBA{R: 233, G: 196, B: 106, A: 255})
	palette = append(palette, color.RGBA{R: 244, G: 162, B: 97, A: 255})
	palette = append(palette, color.RGBA{R: 231, G: 111, B: 81, A: 255})
	palette = append(palette, color.Black)
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
	for k := i * img.colzoom; k < (i+1)*img.colzoom; k++ {
		for l := j * img.rowzoom; l < (j+1)*img.rowzoom; l++ {
			img.raw.Set(k, l, img.raw.Palette[value])
		}
	}
}

func (img *Image) WritePNGToFile() {
	outputFile, _ := os.Create(img.filename)
	png.Encode(outputFile, img.raw)
	outputFile.Close()
}
