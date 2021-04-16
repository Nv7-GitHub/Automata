package main

import (
	"image/png"

	r "github.com/lachee/raylib-goplus/raylib"
	"github.com/nfnt/resize"
)

func initAutomata() {
	var w, h int
	if isHidpi {
		w = width * 2
		h = height * 2
	} else {
		w = width
		h = height
	}

	frame = make([][]Cell, h)
	for i := range frame {
		frame[i] = make([]Cell, w)
	}

	f, err := fs.Open(mapName)
	handle(err)
	img, err := png.Decode(f)
	handle(err)
	img = resize.Resize(uint(len(frame[0])), uint(len(frame)), img, resize.NearestNeighbor)
	handle(err)

	for y := range frame {
		for x := range frame[y] {
			rc, g, b, a := img.At(x, y).RGBA()
			frame[y][x].Color = r.NewColor(uint8(rc), uint8(g), uint8(b), uint8(a))
		}
	}
}
