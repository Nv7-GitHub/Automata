package main

import r "github.com/lachee/raylib-goplus/raylib"

func initAutomata() {
	if isHidpi {
		frame = r.LoadRenderTexture(width*2, height*2)
	} else {
		frame = r.LoadRenderTexture(width, height)
	}
	shader = r.LoadShaderCode(vs, fs)
}
