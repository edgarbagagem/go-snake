package main

import (
	"log"

	"github.com/gdamore/tcell"
)

func main(){
	screen, err := tcell.NewScreen();

	if err != nil {
		log.Fatalf("Failed to create screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Failed to initialize screen: %v", err)
	}

	defer screen.Fini();

	//Initial Game Settings
	initialBody := []BodyPart{
		{
			X: 5,
			Y: 10,
		},
		{
			X: 6,
			Y: 10,
		},
		{
			X: 7,
			Y: 10,
		},
		{
			X: 8,
			Y: 10,
		},
		{
			X: 9,
			Y: 10,
		},
	}

	snake := Snake{

		Xspeed: 1,
		Yspeed: 0,
		Body: initialBody,
		
	}

	game := Game{
		Screen:	screen,
		Snake: snake,
		Paused: false,
	}

	go game.Run()

	//Game Funcionalities
	for {

		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync();
		case *tcell.EventKey: 
			if(event.Key() == tcell.KeyEscape){ // Exit game
				return;
			}
			if(event.Key() == tcell.KeyUp){
				game.Snake.ChangeDir(-1, 0);
			}
			if(event.Key() == tcell.KeyDown){
				game.Snake.ChangeDir(1, 0);
			}
			if(event.Key() == tcell.KeyLeft){
				game.Snake.ChangeDir(0 , -1);
			}
			if(event.Key() == tcell.KeyRight){
				game.Snake.ChangeDir(0, 1);
			}
			if(event.Rune() == 'p'){
				game.Paused = !game.Paused
			}
		}

	}
	
}