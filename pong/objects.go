package pong

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

//GameState determines what the game does depending on it's current state
type GameState int

const (
	//SETUP preps the game state
	SETUP GameState = iota
	//Play is the game state while playing
	Play
	//Stop is the game state while not playing
	Stop
)

//Paddle defines the attributes of the paddle entity
type Paddle struct {
	Width  int
	Length int
	PosX   float64
	PosY   float64
	VelX   float64
	VelY   float64
}

const (
	//PaddleWidth is the width of the paddles
	PaddleWidth = 10
	//PaddleLength is the length of the paddles
	PaddleLength = 30
	//PaddleVelY is the speed of the paddles on the Y axis
	PaddleVelY = 2
	//PaddleVelX is the speed of the paddles on the X axis
	PaddleVelX = 2
	//PaddlePosX is the staring postion of the paddles on the X axis
	PaddlePosX = 0
	//PaddlePosY is the staring postion of the paddles on the Y axis
	PaddlePosY = 15
)

var (
	drawing ebiten.Image
)

//Ball defines the attributes of the ball entity
type Ball struct {
	Size   int
	AccelX int
	AccelY int
	VelX   float64
	VelY   float64
	PosX   float64
	PosY   float64
	InPlay GameState
}

const (
	//BallSize is the balls size duh
	BallSize = 10
	//BallAccelX is the accelaration factor on the X axis
	BallAccelX = 1
	//BallAccelY is the accelaration factor on the Y axis
	BallAccelY = 1
	//BallPosX is dictates the staring position of the ball on the X axis
	BallPosX = 50
	//BallPosY is dictates the staring position of the ball on the Y axis
	BallPosY = 50
	//BallVelX determines the starting velocity of the ball across the X axis
	BallVelX = 2
	//BallVelY determines the starting velocity of the ball across the Y axis
	BallVelY = 2
)

func (b *Ball) initBall(ballSPosX float64, ballSPosY float64) {
	b.Size = BallSize
	b.AccelX = BallAccelX
	b.AccelY = BallAccelY
	b.PosX = BallPosX - float64(b.Size*2)
	b.PosY = BallPosY - float64(b.Size*2)
	b.InPlay = Play
}
func (p *Paddle) initPaddles(paddleSPosX float64, paddleSPosY float64) {
	p.Length = PaddleLength
	p.Width = PaddleWidth
	p.PosX = PaddlePosX
	p.PosY = PaddlePosY
	p.VelX = PaddleVelX
	p.VelY = PaddleVelY

}

func (p *Paddle) drawPaddle(screen *ebiten.Image, paddleSPosX float64, paddleSPosY float64) error {
	var err error

	if Play == SETUP {
		p.initPaddles(paddleSPosX, paddleSPosY)

	}

	screen, err = ebiten.NewImage(p.Width, p.Length, ebiten.FilterNearest)
	if err != nil {
		log.Print(err)
	}

	if err = screen.Fill(color.White); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (b *Ball) drawBall(screen *ebiten.Image, ballSPosX float64, ballSPosY float64) error {
	var err error

	if b.InPlay == SETUP {
		b.initBall(ballSPosX, ballSPosY)
	}

	screen, err = ebiten.NewImage(2*b.Size, 2*b.Size, ebiten.FilterNearest)
	if err != nil {
		log.Print(err)
		return err
	}

	if err = screen.Fill(color.White); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (b *Ball) bPlay() {
	if b.InPlay != Play {
		b.InPlay = Play
	}
}

func (b *Ball) ballPlaying(screen *ebiten.Image) {
	var screenHeight, screenWidth int
	screenHeight, screenWidth = screen.Size()

	b.PosX = b.PosX + b.VelX
	b.PosY = b.PosY + b.VelY

	ballMovement := &ebiten.DrawImageOptions{}
	ballMovement.GeoM.Translate(b.PosX, b.PosY)

	if b.PosX == 0 {
		b.VelX = b.VelX * -1
	}
	if b.PosY == 0 {
		b.VelY = b.VelY * -1
	}

	if b.PosY == float64(screenHeight-(b.Size*2)) {
		b.VelY = b.VelY * -1
	}

	if b.PosX == float64(screenWidth-(b.Size*2)) {
		b.VelX = b.VelX * -1
	}

}

func (b *Ball) ballStopping(screen *ebiten.Image) {
	b.PosX = BallPosX
	b.PosY = BallPosY

	ballMovement := &ebiten.DrawImageOptions{}
	ballMovement.GeoM.Translate(b.PosX, b.PosY)

}

func (p *Paddle) paddlePlay(screen *ebiten.Image) {

}
