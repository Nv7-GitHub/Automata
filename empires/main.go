package main

import (
	"image"

	r "github.com/lachee/raylib-goplus/raylib"
)

var colorIndex = 0

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r.InitWindow(width, height, name)
	r.SetTargetFPS(fps)
	defer r.UnloadAll()
	initAutomata()
	r.SetTraceLogLevel(r.LogWarning | r.LogError | r.LogFatal | r.LogTrace)

	// Benching
	/*out, err := os.Create("prof.pprof")
	handle(err)
	err = pprof.StartCPUProfile(out)
	handle(err)
	defer pprof.StopCPUProfile()*/

	frm := 0

	for !r.WindowShouldClose() {
		if isHidpi {
			r.SetMouseScale(2, 2)
		}

		if frm == fpsim {
			frm = 0
			simulate()
		}

		if r.IsMouseButtonDown(r.MouseLeftButton) {
			pos := r.GetMousePosition()
			min := image.Pt(int(pos.X)-mouseWidth/2, int(pos.Y)-mouseHeight/2)
			max := image.Pt(mouseWidth, mouseHeight).Add(min)
			for y := min.Y; y < max.Y; y++ {
				for x := min.X; x < max.X; x++ {
					if inBounds(x, y) {
						frame[y][x] = Cell{
							Color:    colors[colorIndex],
							Strength: startingStrength,
						}
					}
				}
			}
		}

		if r.IsKeyPressed(r.KeyRight) && colorIndex < (len(colors)-1) {
			colorIndex++
		} else if r.IsKeyPressed(r.KeyLeft) && colorIndex > 0 {
			colorIndex--
		}

		// Draw
		r.BeginDrawing()

		r.ClearBackground(backgroundColor)

		r.DrawTextureEx(tex, r.NewVector2(0, 0), 0, 0.5, r.White)

		r.DrawFPS(20, height-40)
		r.DrawText("Color", width/2-50, 20, 20, colors[colorIndex])

		r.EndDrawing()

		frm++
	}

	r.CloseWindow()
}
