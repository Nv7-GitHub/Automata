package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

var out [][]Cell

func simulate() {
	out = make([][]Cell, len(frame))
	for i := range out {
		out[i] = make([]Cell, len(frame[i]))
	}

	for y := 0; y < len(out); y++ {
		for x := 0; x < len(frame[y]); x++ {
			simulatePos(x, y)
		}
	}

	frame = out

	tex.Unload()
	tex = getFrameTex(frame)
}

func getFrameTex(frame [][]Cell) r.Texture2D {
	pixels := make([]r.Color, len(frame)*len(frame[0]))
	for y, row := range frame {
		for x, cell := range row {
			pixels[y*len(frame[y])+x] = cell.Color
		}
	}
	im := r.LoadImageEx(pixels, int32(len(frame[0])), int32(len(frame)))
	tex := r.LoadTextureFromImage(im)
	im.Unload()
	return tex
}
