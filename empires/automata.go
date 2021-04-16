package main

import (
	"math/rand"

	r "github.com/lachee/raylib-goplus/raylib"
)

func simulatePos(x, y int) {
	// Is player?
	if out[y][x].Color != r.Blank {
		return
	}
	if (frame[y][x].Color == groundColor) || (frame[y][x].Color == waterColor) {
		out[y][x] = frame[y][x]
		return
	}

	// Aging
	frame[y][x].Age++
	frame[y][x].ReproductionValue++
	if frame[y][x].Age > frame[y][x].Strength {
		out[y][x].Color = groundColor
		frame[y][x].Color = groundColor
		return
	}

	// Reproduction
	if frame[y][x].ReproductionValue > reproductionThreshold {
		var newX, newY int
		foundPos := false
		for x2 := -1; x2 <= 1; x2++ {
			for y2 := -1; y2 <= 1; y2++ {
				if inBounds(x+x2, y+y2) && frame[y+y2][x+x2].Color == groundColor {
					newX = x + x2
					newY = y + y2
					foundPos = true
					break
				}
			}
		}

		if foundPos {
			curr := frame[y][x]
			child := Cell{
				Color:    curr.Color,
				Strength: curr.Strength,
			}
			if curr.Diseased {
				child.Diseased = rand.Intn(100) >= 85
			}

			mutation := rand.Intn(1000)
			if mutation >= 999 {
				child.Diseased = true
				child.Strength = (child.Strength * 13) / 20 // 65% of original strength
			} else if mutation >= 750 {
				child.Strength = int(float64(child.Strength) * rand.Float64())
			}

			frame[newY][newX] = child
			out[newY][newX] = child
		}
	}

	// Movement
	newX := x + rand.Intn(3) - 1
	newY := y + rand.Intn(3) - 1
	for !inBounds(newX, newY) {
		newX = rand.Intn(3) - 1
		newY = rand.Intn(3) - 1
	}
	if frame[newY][newX].Color == groundColor {
		if out[newY][newX].Color != r.Blank {
			out[y][x] = out[newY][newX]
		} else {
			out[y][x].Color = groundColor
		}
		out[newY][newX] = frame[y][x]
	} else {
		out[y][x] = frame[y][x]
	}
}

func inBounds(x, y int) bool {
	return x >= 0 && y >= 0 && y < len(frame) && x < len(frame[y])
}
