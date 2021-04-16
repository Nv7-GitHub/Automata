package main

import (
	_ "embed"
	"image/color"

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

type Cell struct {
	Color    r.Color
	Age      int
	Strength int
}

var frame [][]Cell
var tex r.Texture2D
