package main

type BodyPart struct {
	X int
	Y int
}

type Snake struct {
	Xspeed int
	Yspeed int
	Body   []BodyPart
}

func (snake *Snake) ChangeDir(vertical int, horizontal int) {
	snake.Yspeed = vertical
	snake.Xspeed = horizontal
}

func (snake *Snake) Update(width int, height int) {
	snake.Body = append(snake.Body, snake.Body[len(snake.Body)-1].MovePart(snake, width, height))
	snake.Body = snake.Body[1:]
}

func (bodyPart *BodyPart) MovePart(snake *Snake, width int, height int) BodyPart {
	newBodyPart := *bodyPart
	newBodyPart.X = (newBodyPart.X + snake.Xspeed) % width
	if newBodyPart.X < 0 {
		newBodyPart.X += width

	}

	newBodyPart.Y = (newBodyPart.Y + snake.Yspeed) % width
	if newBodyPart.Y < 0 {
		newBodyPart.Y += height
	}

	return newBodyPart
}