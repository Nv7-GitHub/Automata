package main

import (
	_ "embed"

	r "github.com/lachee/raylib-goplus/raylib"
)

const (
	width   = 800
	height  = 450
	isHidpi = true
	name    = "Empires"

	mouseWidth  = 50
	mouseHeight = 50

	fps = 60
)

var (
	backgroundColor = r.RayWhite

	conducterColor = r.Gray
	headColor      = r.Blue
	tailColor      = r.Red
	colors         = []Tool{
		{
			Name:  "Conducter",
			Color: conducterColor,
		},
		{
			Name:  "Head",
			Color: headColor,
		},
		{
			Name:  "Tail",
			Color: tailColor,
		},
	}
)

type Tool struct {
	Name  string
	Color r.Color
}

//go:embed default.vs
var vs string

//go:embed automata.fs
var fs string

var frame r.RenderTexture2D
var shader r.Shader
