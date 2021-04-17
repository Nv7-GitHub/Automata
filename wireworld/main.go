package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

var colorIndex = 0

func main() {
	r.InitWindow(width, height, name)
	r.SetTargetFPS(fps)
	defer r.UnloadAll()
	initAutomata()
	r.SetTraceLogLevel(r.LogDebug | r.LogError | r.LogFatal | r.LogWarning)

	for !r.WindowShouldClose() {
		if isHidpi {
			r.SetMouseScale(2, 2)
		}

		// Basic Drawing
		r.BeginDrawing()

		out := getRenderTexture()

		r.BeginTextureMode(out)
		r.BeginShaderMode(shader)

		r.ClearBackground(r.Blank)

		r.DrawTextureRec(frame.Texture, r.NewRectangle(0, 0, float32(frame.Texture.Width), float32(-frame.Texture.Height)), r.NewVector2(0, 0), r.White)

		// Input
		if r.IsKeyPressed(r.KeyRight) && colorIndex < (len(colors)-1) {
			colorIndex++
		} else if r.IsKeyPressed(r.KeyLeft) && colorIndex > 0 {
			colorIndex--
		}
		if r.IsMouseButtonDown(r.MouseLeftButton) {
			pos := r.GetMousePosition()
			r.DrawRectangle(int(pos.X)-mouseWidth/2, int(pos.Y)-mouseHeight/2, mouseWidth, mouseHeight, colors[colorIndex].Color)
		}

		r.EndShaderMode()
		r.EndTextureMode()

		frame.Unload()
		frame = out

		// Draw
		r.ClearBackground(backgroundColor)
		r.DrawTextureRec(frame.Texture, r.NewRectangle(0, 0, float32(frame.Texture.Width), float32(-frame.Texture.Height)), r.NewVector2(0, 0), r.White)
		toolName := colors[colorIndex].Name
		if !isHidpi {
			r.DrawFPS(20, height-20)
			r.DrawText(toolName, width-(len(toolName)), 20, 20, colors[colorIndex].Color)
		} else {
			r.DrawFPS(20, height*2-40)
			r.DrawText(toolName, width-(len(toolName)*10), 20, 20, colors[colorIndex].Color)
		}

		r.EndDrawing()
	}

	r.CloseWindow()
}
