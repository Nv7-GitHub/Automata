package main

import (
	_ "embed"
	"image"
	"image/color"
	"image/draw"

	r "github.com/lachee/raylib-goplus/raylib"
)

const (
	width   = 800
	height  = 450
	isHidpi = true
	name    = "Empires"

	mouseWidth  = 50
	mouseHeight = 50

	fps   = 60
	fpsim = 1
)

var (
	backgroundColor = r.RayWhite
	mouseColor      = color.RGBA{
		R: 0,
		G: 200,
		B: 255,
		A: 255,
	}
)

var frame draw.Image
var rect image.Rectangle
var tex r.Texture2D
