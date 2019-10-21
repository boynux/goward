package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
  "time"

	"github.com/gen2brain/raylib-go/raygui"
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	TotalQuestionsPerScene = 30
  ScreenSaverTimeout = 60 * time.Second
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	exit := false

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	screenWidth := int32(280)
	screenHeight := int32(210)

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	raygui.LoadGuiStyle("styles/solarized_light.style")

	bg := NewBasicGenerator(0, 20, nil, []string{"+", "-"})
	ag := NewBasicAdditionGenerator(10, 20)
	even := NewEvenOddGenerator(1, 50)

	s := NewScenario([]Question{NewChoiceQuestion(bg), NewChoiceQuestion(even), NewRectangleChoiceQuestion(ag)}, TotalQuestionsPerScene, 1)
	s.Order(Random)

  activityCheck := time.NewTimer(ScreenSaverTimeout)
  saveMode := false
  var animation *Animation

	for !exit && !rl.WindowShouldClose() {
    if rl.IsMouseButtonReleased(rl.MouseLeftButton) || rl.IsMouseButtonDown(rl.MouseLeftButton) {
      activityCheck = time.NewTimer(ScreenSaverTimeout)
      saveMode = false
    }

		rl.BeginDrawing()

    if !saveMode {
      rl.ClearBackground(raygui.BackgroundColor())

      r, c := s.Repeats()

      raygui.Label(rl.NewRectangle(float32(screenWidth-80), 5, 75, 20), fmt.Sprintf("Correct: %d", c))

      if s.Play() == false {
        if showResults(r, c) {
          s.Restart()
        }
      }

      showProgress(TotalQuestionsPerScene, r)
      exit = raygui.Button(rl.NewRectangle(float32(screenWidth-60-5), float32(screenHeight-30-5), 60, 30), "Exit")
    } else {
      rl.ClearBackground(rl.Black)
      drawPolygon(animation)
    }

		rl.EndDrawing()

		select {
		case <-done:
			exit = true
    case <-activityCheck.C:
      animation = NewAnimation()
      saveMode = true
		default:
		}
	}

	rl.CloseWindow()
}

func showProgress(total, progress int32) {
	p := float32(progress) / float32(total)
	r := rl.NewRectangle(5, float32(rl.GetScreenHeight()-30-5), float32(rl.GetScreenWidth()-60-5-5-5), 30)

	raygui.ProgressBar(r, p)
	raygui.Label(rl.NewRectangle(float32(rl.GetScreenWidth())/2-35, float32(rl.GetScreenHeight()-30), 20, 20), fmt.Sprintf("%d", progress+1))
}


var anim *Animation
var ticker *time.Ticker

func drawPolygon(a *Animation) {
  if ticker == nil {
    ticker = time.NewTicker(20 * time.Millisecond)
  }

  a.Draw()

  select{
  case <-ticker.C:
    a.Update()
  default:
  }
}
