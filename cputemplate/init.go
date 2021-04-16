package main

import (
	"image"

	r "github.com/lachee/raylib-goplus/raylib"
)

func initAutomata() {
	if isHidpi {
		rect = image.Rect(0, 0, width*2, height*2)
	} else {
		image.Rect(0, 0, width, height)
	}
	frame = image.NewRGBA(rect)
	tex = r.LoadTextureFromGo(frame)
}
