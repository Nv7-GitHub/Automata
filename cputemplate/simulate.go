package main

import (
	"image"
	"image/draw"

	r "github.com/lachee/raylib-goplus/raylib"
)

func simulate() {
	tex.Unload()
	a := &Automata{
		frame,
	}
	draw.Draw(frame, frame.Bounds(), a, image.Point{}, draw.Src)
	tex = r.LoadTextureFromGo(frame)
}
