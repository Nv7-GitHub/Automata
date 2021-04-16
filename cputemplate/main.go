package main

import (
	"image"

	r "github.com/lachee/raylib-goplus/raylib"
)

func main() {
	r.InitWindow(width, height, name)
	r.SetTargetFPS(fps)
	defer r.UnloadAll()
	initAutomata()
	r.SetTraceLogLevel(r.LogWarning | r.LogError | r.LogFatal | r.LogTrace)

	frm := 0

	for !r.WindowShouldClose() {
		if isHidpi {
			r.SetMouseScale(1, 1)
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
					frame.Set(x, y, mouseColor)
				}
			}
			simulate()
		}

		// Draw
		r.BeginDrawing()

		r.ClearBackground(backgroundColor)
		r.DrawTexture(tex, 0, 0, r.White)
		r.DrawFPS(20, height-40)

		r.EndDrawing()

		frm++
	}

	r.CloseWindow()
}
