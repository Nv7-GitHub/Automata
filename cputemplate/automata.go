package main

import (
	"image/color"
	"image/draw"
	"math/rand"
)

type Automata struct {
	draw.Image
}

func (a *Automata) At(x, y int) color.Color {
	offX := rand.Intn(3) - 1
	offY := rand.Intn(3) - 1
	color := a.Image.At(x+offX, y+offY)
	return color
}
