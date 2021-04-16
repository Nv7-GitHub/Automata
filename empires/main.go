package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

func main() {
	r.InitWindow(width, height, name)
	r.SetTargetFPS(fps)
	defer r.UnloadAll()
	initAutomata()

	for !r.WindowShouldClose() {
		if isHidpi {
			r.SetMouseScale(2, 2)
		}

		// Basic Drawing
		r.BeginDrawing()
		r.ClearBackground(r.Blank)

		r.BeginShaderMode(shader)
		r.BeginTextureMode(frame)

		if r.IsMouseButtonPressed(r.MouseLeftButton) {
			pos := r.GetMousePosition()
			r.DrawRectangle(int(pos.X)-mouseWidth/2, int(pos.Y)-mouseHeight/2, mouseWidth, mouseHeight, r.Blue)
		}

		r.EndTextureMode()
		r.EndShaderMode()

		// Draw
		r.ClearBackground(backgroundColor)
		r.DrawTextureRec(frame.Texture, r.NewRectangle(0, 0, float32(frame.Texture.Width), float32(-frame.Texture.Height)), r.NewVector2(0, 0), r.White)
		if !isHidpi {
			r.DrawFPS(20, height-20)
		} else {
			r.DrawFPS(20, height*2-40)
		}

		r.EndDrawing()
	}

	r.CloseWindow()
}
