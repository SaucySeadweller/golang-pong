package pong

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type currentScore struct {
	Score int
	x     float64
	y     float64
	print bool
}

//Pos is the positions x & y
type Pos struct {
	X, Y float64
}

var (
	//Background color
	Background = color.Black
	//PaddleBall paddle and ball color
	PaddleBall = color.White
)

type keyPressed struct {
	W bool
	S bool
}

//Paddle defines the attributes of the paddle entity
type Paddle struct {
	Pos
	Width        int
	Length       int
	Velocity     float64
	Score        int
	CurrentScore currentScore
	Paddle       *ebiten.Image
	Speed        float64
	KeyPressed   keyPressed
	PaddleColor  color.Color
	MovementUp   ebiten.Key
	MovementDown ebiten.Key
}

const (
	//PaddleWidth is the width of the paddles
	PaddleWidth = 10
	//PaddleLength is the length of the paddles
	PaddleLength = 100
	//PaddleStart is the paddle start position
	PaddleStart = 45
)

//Update ypdates the paddles movements on screen
func (p *Paddle) Update(screen *ebiten.Image) {
	_, screenHeight := screen.Size()

	if inpututil.IsKeyJustPressed(p.MovementUp) {
		p.KeyPressed.S = false
		p.KeyPressed.W = true
	} else if inpututil.IsKeyJustReleased(p.MovementUp) || !ebiten.IsKeyPressed(p.MovementUp) {
		p.KeyPressed.W = false
	}
	if inpututil.IsKeyJustPressed(p.MovementDown) {
		p.KeyPressed.W = false
		p.KeyPressed.S = true
	} else if inpututil.IsKeyJustReleased(p.MovementDown) || !ebiten.IsKeyPressed(p.MovementDown) {
		p.KeyPressed.S = false
	}

	if p.KeyPressed.W {
		p.Y -= p.Speed
	} else if p.KeyPressed.S {
		p.Y += p.Speed
	}

	if p.Y-float64(p.Length/2) < 0 {
		p.Y = float64(1 + p.Length/2)
	} else if p.Y+float64(p.Length/2) > float64(screenHeight) {
		p.Y = float64(screenHeight - p.Length/2 - 1)
	}
}

//PaddleControls defines the InPlay actions of the paddles
func (p *Paddle) PaddleControls(screen *ebiten.Image) {
	var err error
	Opts := &ebiten.DrawImageOptions{}
	Opts.GeoM.Translate(p.X, p.Y)
	if err = screen.DrawImage(screen, Opts); err != nil {
		log.Print(err)
	}
}

//AI that is hard to beat
func (p *Paddle) AI(b *Ball) {
	p.Y = b.X
}
