package pong

import (
	"image/color"
	"log"
	"strconv"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
)

type currentScore struct {
	score int
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
		p.Y = float64(-p.Length/2 - 1)
	}
}

//Draw draws the paddle
func (p *Paddle) Draw(screen *ebiten.Image, f font.Face) {
	paddle := &ebiten.DrawImageOptions{}
	paddle.GeoM.Translate(p.X, p.Y-PaddleLength/2)
	p.Paddle.Fill(color.White)
	screen.DrawImage(p.Paddle, paddle)

	s := strconv.Itoa(p.CurrentScore.score)
	text.Draw(screen, s, f, int(p.CurrentScore.x), int(p.CurrentScore.y), color.White)
	if p.CurrentScore.score != p.Score && p.CurrentScore.print {
		p.CurrentScore.print = false
	}
	if p.CurrentScore.score == 0 && !p.CurrentScore.print {
		p.CurrentScore.x = (p.X + (Center(screen).X-p.X)/2)
		p.CurrentScore.y = (2 * 20)
	}
	if (p.CurrentScore.score == 0 || p.CurrentScore.score != p.Score) && !p.CurrentScore.print {
		p.CurrentScore.score = p.Score
		p.CurrentScore.print = true
	}

}

//PaddleControls defines the InPlay actions of the paddles
func (p *Paddle) PaddleControls(screen *ebiten.Image) {
	var err error
	collision := &ebiten.DrawImageOptions{}
	collision.GeoM.Translate(p.X, p.Y)
	if err = screen.DrawImage(draw, collision); err != nil {
		log.Print(err)
	}
}

//AI does the thing
func (p *Paddle) AI(b *Ball) {
	p.Y = b.Y
}