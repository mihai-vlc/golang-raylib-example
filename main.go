package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	var fontSize float32 = 80
	var target float32 = 120.0
	for !rl.WindowShouldClose() {
		var dt = rl.GetFrameTime()
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(0x18, 0x18, 0x18, 255))

		fontSize = rl.Lerp(fontSize, target, 0.6*dt)

		if fontSize > 100 {
			target = 40
		}

		if fontSize < 80 {
			target = 120
		}

		// DrawCenteredText("Hello darkness my old friend !", 20, rl.White)
		DrawCenteredText("Welcome to my app", int32(fontSize), rl.White)

		rl.EndDrawing()
	}
}

func DrawCenteredText(text string, fontSize int32, color rl.Color) {
	var width = rl.GetRenderWidth()
	var height = rl.GetRenderHeight()
	var textSize = rl.MeasureText(text, fontSize)
	rl.DrawText(text, int32(width)/2-textSize/2, int32(height)/2-fontSize/2, fontSize, color)
}
