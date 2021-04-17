package main

import r "github.com/lachee/raylib-goplus/raylib"

func initAutomata() {
	frame = getRenderTexture()
	shader = r.LoadShaderCode(vs, fs)
}

func getRenderTexture() r.RenderTexture2D {
	if isHidpi {
		return r.LoadRenderTexture(width*2, height*2)
	}
	return r.LoadRenderTexture(width, height)
}
