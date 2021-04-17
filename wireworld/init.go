package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

func initAutomata() {
	frame = getRenderTexture()
	shader = r.LoadShaderCode(vs, fs)
	shader.SetValueFloat32(shader.GetLocation("conducterColor"), conducterColor.Normalize().Decompose(), r.UniformVec4)
	shader.SetValueFloat32(shader.GetLocation("headColor"), headColor.Normalize().Decompose(), r.UniformVec4)
	shader.SetValueFloat32(shader.GetLocation("tailColor"), tailColor.Normalize().Decompose(), r.UniformVec4)
}

func getRenderTexture() r.RenderTexture2D {
	if isHidpi {
		return r.LoadRenderTexture(width*2, height*2)
	}
	return r.LoadRenderTexture(width, height)
}
