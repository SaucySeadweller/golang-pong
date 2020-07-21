package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var screen *ebiten.Image

const (
	screenWidth  = 640
	screenHeight = 480
	gamestate    = iota
)

func main() {
	var err error
	screen, err = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterNearest)
	if err != nil {
		log.Print(err)
	}
	errTrue := ebiten.Run(drawScreen, screenWidth, screenHeight, 1, "Pong")
	if errTrue != nil {
		log.Print(err)
	}
}

func drawScreen(screen *ebiten.Image) error {
	var err error
	screen.Fill(color.Black)

	if err != nil {
		log.Print(err)
	}

	return err
}
