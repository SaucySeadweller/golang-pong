package pong

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

//GameState determines what the game does depending on it's current state
type GameState byte

//Ball defines the attributes of the ball entity
type Ball struct {
	Pos
	Radius float64
	AccelX int
	AccelY int
	VelX   float64
	VelY   float64
	Ball   *ebiten.Image
	Color  color.Color
}

const (
	//BallRadius is the balls Radius duh
	BallRadius = 10.0
	//BallVelX determines the starting velocity of the ball across the X axis
	BallVelX = 2
	//BallVelY determines the starting velocity of the ball across the Y axis
	BallVelY = 2
	//BallStart is the balls starting position  on screen
	BallStart = 50
)

//Update updates the acions on screen
func (b *Ball) Update(screen *ebiten.Image, playerPaddle *Paddle, aiPaddle *Paddle) {
	_, screenHeight := screen.Size()

	b.X += b.VelX
	b.Y += b.VelY

	if b.Y-b.Radius > float64(screenHeight) {
		b.VelY = -b.VelY
		b.Y = float64(screenHeight) - b.Radius
	} else if b.Y+b.Radius < 0 {
		b.VelY = -b.VelY
		b.Y = b.Radius
	}

	if b.X-b.Radius < playerPaddle.X+float64(playerPaddle.Width/2) &&
		b.Y > playerPaddle.Y-float64(playerPaddle.Length/2) &&
		b.Y < playerPaddle.Y+float64(playerPaddle.Length/2) {
		b.VelX = -b.VelX
		b.X = playerPaddle.X + float64(playerPaddle.Width/2) + b.Radius
	} else if b.X+b.Radius > aiPaddle.X-float64(aiPaddle.Width/2) &&
		b.Y > aiPaddle.Y-float64(aiPaddle.Length/2) &&
		b.Y < aiPaddle.Y+float64(aiPaddle.Length/2) {
		b.VelX = -b.VelX
		b.X = aiPaddle.X - float64(aiPaddle.Width/2) - b.Radius
	}

}
