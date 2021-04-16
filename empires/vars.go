package main

import (
	"embed"
	"image/color"

	r "github.com/lachee/raylib-goplus/raylib"
)

//go:embed map.png
var fs embed.FS

const (
	mapName = "map.png"

	width   = 800
	height  = 450
	isHidpi = true
	name    = "Empires"

	mouseWidth  = 50
	mouseHeight = 50

	fps   = 60
	fpsim = 1

	startingStrength      = 100
	reproductionThreshold = 50
)

var (
	backgroundColor = r.RayWhite
	mouseColor      = color.RGBA{
		R: 0,
		G: 200,
		B: 255,
		A: 255,
	}
	groundColor = r.NewColor(0, 255, 0, 255)
	waterColor  = r.NewColor(0, 0, 255, 255)
)

type Cell struct {
	Color             r.Color
	Age               int
	Strength          int
	ReproductionValue int
	Diseased          bool
}

var frame [][]Cell
var tex r.Texture2D
