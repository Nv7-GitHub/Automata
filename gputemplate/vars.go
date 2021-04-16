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
)

//go:embed default.vs
var vs string

//go:embed automata.fs
var fs string

var frame r.RenderTexture2D
var shader r.Shader
