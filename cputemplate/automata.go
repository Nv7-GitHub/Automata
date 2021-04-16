package main

import (
	"math/rand"

	r "github.com/lachee/raylib-goplus/raylib"
)

func simulatePos(x, y int) {
	if frame[y][x].Color == r.Blank {
		return
	}
	newX := x + rand.Intn(3) - 1
	newY := y + rand.Intn(3) - 1
	for !inBounds(newX, newY) {
		newX = rand.Intn(3) - 1
		newY = rand.Intn(3) - 1
	}
	out[newY][newX] = frame[y][x]
}

func inBounds(x, y int) bool {
	return x >= 0 && y >= 0 && y < len(frame) && x < len(frame[y])
}
