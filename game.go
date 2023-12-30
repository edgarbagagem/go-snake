package main

import (
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen tcell.Screen
	Snake Snake
	Paused bool
}

func drawSnake(screen tcell.Screen, body []BodyPart, style tcell.Style){
	for _, bodyPart := range body {
		screen.SetContent(bodyPart.X, bodyPart.Y, '-', []rune{'|'}, style)
	}
}

func (game *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack)
	game.Screen.SetStyle(defStyle)
	width, height := game.Screen.Size()
	snakeColor := tcell.StyleDefault.Background(tcell.ColorDarkGreen).Foreground(tcell.ColorYellow);

	for{
		if(game.Paused == true) {
			continue
		}
		game.Screen.Clear()
		game.Snake.Update(width, height);
		drawSnake(game.Screen, game.Snake.Body, snakeColor)
		time.Sleep(40 * time.Millisecond)
		game.Screen.Show()
	}
} 