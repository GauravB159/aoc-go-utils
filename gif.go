package aocutils

import (
	"image"
	"image/gif"
	"os"
)

type GIF struct {
	Images     []*image.Paletted
	Filename   string
	Frameskip  int
	Framecount int
}

func CreateGIF(filename string, frameskip int) GIF {
	return GIF{
		Images:     make([]*image.Paletted, 0),
		Filename:   filename + ".gif",
		Frameskip:  frameskip,
		Framecount: 0,
	}
}

func (g *GIF) AddFrame(image Image) {
	g.Framecount += 1
	if g.Framecount%g.Frameskip == 0 {
		copy := image.Clone()
		g.Images = append(g.Images, copy.GetRawImage())
	}
}

func (g *GIF) WriteGIFToFile() {
	var delays []int
	for range g.Images {
		delays = append(delays, 1)
	}

	lastFrame := g.Images[len(g.Images)-1]
	for i := 0; i < g.Frameskip*2; i++ {
		g.Images = append(g.Images, lastFrame)
		delays = append(delays, 100)
	}
	gifFile, err := os.Create(g.Filename)
	if err != nil {
		panic(err)
	}
	defer gifFile.Close()

	err = gif.EncodeAll(gifFile, &gif.GIF{
		Image: g.Images,
		Delay: delays,
	})
	if err != nil {
		panic(err)
	}
}
