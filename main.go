package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

var screen *ebiten.Image

const (
	screenWidth  = 640
	screenHeight = 480
	gamestate    = iota
)

type game struct {
	player *Player
	state  byte
	ball   *ball
	pong.Paddle *paddle
}

//Player defines the attributes of the player
type Player struct {
	prevX  int
	prevY  int
	score  int
	paddle *paddle
}

type Physical interface {
	Position() (int, int)
	Size() (int, int)
}

type DynamicPhysical interface {
	Position(int, int)
	Size() (int, int)
	Collide(Physical)
}

func (game *game) initObjects() {
	game.player.Paddle.length = paddleLength
	game.player.Paddle.width = paddleWidth
	game.player.Paddle.posX = paddlePosX
	game.player.Paddle.posY = paddlePosY
	game.player.Paddle.velX = paddleVelX
	game.player.Paddle.velY = paddleVelY
	game.Ball.size = ballSize
	game.Ball.accelX = ballAccelX
	game.Ball.accelY = ballAccelY
	game.Ball.posX = ballPosX
	game.Ball.posY = ballPosY

}

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

//Tick gives the player actions
func (game *game) Tick(key ebiten.Key) {

	switch key {
	case ebiten.KeyUp:
		game.player.paddle.posY = game.player.paddle.posY - game.player.paddle.velY
	case ebiten.KeyDown:
		game.player.paddle.posY = game.player.paddle.posY + game.player.paddle.velY
	}

}
func score() {

}

func random() int {
	min := 1
	max := 2
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func (game *game) winner() bool {
	var score int
	if score == 5 {
		fmt.Println(game.player, "wins!")
		return true
	}

	return false

}

func drawScreen(screen *ebiten.Image) error {
	var err error
	screen.Fill(color.Black)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (game *game) drawBall(screen *ebiten.Image, posX float64, posY float64) {

}

func (game *game) drawPaddle(screen *ebiten.Image, posX float64, posY float64) error {
	var err error
	game.initObjects()
	screen, err = ebiten.NewImage(game.player.paddle.width, game.player.paddle.length, ebiten.FilterNearest)
	if err != nil {
		log.Print(err)
	}

	if err = screen.Fill(color.White); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

//Start game
//Track score
//Paddle movement
//ball movement
//wall,paddle and ball collision
//display score
//finish at determined score
//Boundaries
