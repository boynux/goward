package main

import (
	"fmt"

	"github.com/boynux/goward/actions"
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const Message = "%d correct out of %d questions"
var actionTriggered = false

func showResults(total, correct int32, a actions.Action) bool {

	if float32(correct) / float32(total) * 100.0  < 70 {
		actionTriggered = true
	}

	s := fmt.Sprintf(Message, correct, total)
	o := rl.MeasureText(s, rl.GetFontDefault().BaseSize)
	r := rl.NewRectangle(float32(int32(rl.GetScreenWidth())-o)/2, float32(rl.GetScreenHeight())/2-20, 20, 20)

	raygui.Label(r, s)

	r = rl.NewRectangle(float32(rl.GetScreenWidth())/2-20, float32(rl.GetScreenHeight())/2+30, 40, 20)
	action := raygui.Button(r, "Turn on TV")

	if  !actionTriggered && action {
		a(correct)
	}

	r = rl.NewRectangle(float32(rl.GetScreenWidth())/2-20, float32(rl.GetScreenHeight())/2+5, 40, 20)
	return raygui.Button(r, "Retry")
}
