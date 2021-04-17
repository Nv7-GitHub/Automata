package main

import (
	"math/rand"

	"github.com/Nv7-Github/Automata/util"

	r "github.com/lachee/raylib-goplus/raylib"
)

var bgen = util.NewBoolgen()

func simulatePos(x, y int) {
	me := frame[y][x]
	if ((me.Color == groundColor) || (me.Color == waterColor)) && (out[y][x].Color == r.Blank) {
		out[y][x] = me
		return
	}

	// Is player?
	if out[y][x].Color != r.Blank {
		return
	}

	// Aging
	frame[y][x].Age++
	frame[y][x].ReproductionValue++
	if me.Age > me.Strength {
		out[y][x].Color = groundColor
		frame[y][x].Color = groundColor
		return
	}
	// Aging speeds up when diseased
	if me.Diseased {
		frame[y][x].Age = frame[y][x].Age * 3 / 2
		if bgen.Bool() {
			frame[y][x].Diseased = false // Cured! (50% chance)
		}
	}

	// Reproduction
	if me.ReproductionValue > reproductionThreshold {
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
	newX := x + rand3()
	newY := y + rand3()
	for !inBounds(newX, newY) {
		newX = rand3()
		newY = rand3()
	}
	newCell := frame[newY][newX]
	if newCell.Color == groundColor {
		if out[newY][newX].Color != r.Blank {
			out[y][x] = out[newY][newX]
		} else {
			out[y][x].Color = groundColor
		}
		out[newY][newX] = frame[y][x]
	} else if newCell.Color != waterColor && newCell.Color != me.Color {
		// Fighting
		enemy := frame[newY][newX]
		if me.Strength > enemy.Strength {
			out[newY][newX] = me
		}
		out[y][x].Color = groundColor
	} else {
		out[y][x] = frame[y][x]
		if frame[y][x].Diseased && bgen.Bool() {
			if out[newY][newX].Color == r.Blank {
				out[newY][newX] = frame[newY][newX]
			}
			// Disease spreading
			out[newY][newX].Diseased = true
		}
	}
}

func inBounds(x, y int) bool {
	return x >= 0 && y >= 0 && y < len(frame) && x < len(frame[y])
}

func rand3() int {
	return -1 + bool2int(bgen.Bool()) + bool2int(bgen.Bool())
}

func bool2int(val bool) int {
	if val {
		return 1
	}
	return 0
}
