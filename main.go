package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	maxGestureStrings int = 20
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	exit := false

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	screenWidth := int32(280)
	screenHeight := int32(210)

	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")

	buttonClicked := false
	progressValue := float32(0.0)

	rl.SetTargetFPS(60)

	raygui.LoadGuiStyle("styles/default_dark.style")

	add := NewAdditionGenerator(0, 20, nil)
	s := NewScenario(NewQuestion(add), 10, 2)

	for !buttonClicked && !exit && !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		r, c := s.Repeats()
		progressValue = float32(r) / 10.0 /* it's the total repeats, fix it */

		if s.Play() == false {
			raygui.Label(rl.NewRectangle(float32(screenWidth)/2-80, float32(screenHeight)/2-20, 20, 20), fmt.Sprintf("%d correct out of %d questions", c, r))
		}

		raygui.ProgressBar(rl.NewRectangle(5, float32(screenHeight-30-5), float32(screenWidth-60-5-5-5), 30), progressValue)
		buttonClicked = raygui.Button(rl.NewRectangle(float32(screenWidth-60-5), float32(screenHeight-30-5), 60, 30), "Exit")

		rl.EndDrawing()

		select {
		case <-done:
			exit = true
		default:
		}
	}

	rl.CloseWindow()
}
