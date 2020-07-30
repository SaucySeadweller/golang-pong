package main

import (
	"github.com/SaucySeadweller/golang-pong/pong"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/hajimehoshi/ebiten"
)

//Game is a struct that defines Games attributes
type Game struct {
	state  pong.GameState
	player *pong.Paddle
	ai     *pong.Paddle
	score  int
	ball   *pong.Ball
	speed  int
}

const (
	vel         = 5.0
	paddleSpeed = 10
	update      = 15
	accel       = 0.5
)

var sScreen *ebiten.Image

const (
	screenWidth  = 800
	screenHeight = 600
)

func (g *Game) Layout(width int, height int) (int, int) {

	return height, width
}
func openTheGame() *Game {
	game := &Game{}
	game.init()
	return game
}

func (g *Game) init() {
	g.state = pong.Setup
	g.score = 11
	g.player = &pong.Paddle{
		Pos: pong.Pos{
			X: pong.PaddleStart,
			Y: float64(screenHeight / 2)},
		Score:        0,
		Speed:        paddleSpeed,
		Length:       pong.PaddleLength,
		Width:        pong.PaddleWidth,
		PaddleColor:  pong.PaddleBall,
		MovementUp:   ebiten.KeyW,
		MovementDown: ebiten.KeyS,
	}
	g.ai = &pong.Paddle{
		Pos: pong.Pos{
			X: 469,
			Y: float64(screenHeight / 2)},
		Score:       0,
		Speed:       paddleSpeed,
		Length:      pong.PaddleLength,
		Width:       pong.PaddleWidth,
		PaddleColor: pong.PaddleBall,
	}
	g.ball = &pong.Ball{
		Pos: pong.Pos{

			X: float64(screenWidth / 2),
			Y: float64(screenHeight / 2)},
		VelX:   vel,
		VelY:   vel,
		Radius: pong.BallRadius,
		Color:  pong.PaddleBall,
	}

	g.player.Paddle, _ = ebiten.NewImage(g.player.Width, g.player.Length, ebiten.FilterDefault)
	g.ai.Paddle, _ = ebiten.NewImage(g.ai.Width, g.ai.Length, ebiten.FilterDefault)
	g.ball.Ball, _ = ebiten.NewImage(int(g.ball.Radius)*2, int(g.ball.Radius)*2, ebiten.FilterDefault)
	pong.InitFont()
}

func (g *Game) drawScreen(screen *ebiten.Image) error {
	//	var err error
	screen.Fill(pong.Background)
	pong.BigTextDraw(g.state, pong.PaddleBall, screen)
	pong.TextDraw(g.state, pong.PaddleBall, screen)

	if g.state != pong.Key {
		g.player.Draw(screen, pong.Font)
		g.ai.Draw(screen, pong.Font)
		g.ball.Draw(screen)
	}

	return nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	switch g.state {
	case pong.Setup:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = pong.Play

		} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
			g.state = pong.Key
		}
	case pong.Key:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = pong.Play

		}

	case pong.Play:
		screenWidth, _ := screen.Size()

		g.player.Update(screen)
		g.ai.AI(g.ball)

		g.ai.Update(screen)

		g.ball.Update(screen, g.player, g.ai)

		g.speed++

		if g.ball.X < 0 {
			g.ai.Score++

		} else if g.ball.X > float64(screenWidth) {
			g.player.Score++
			g.restart(screen, pong.Play)
		}
		if g.player.Score == g.score || g.ai.Score == g.score {
			g.restart(screen, pong.Play)
		}

	case pong.GameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.restart(screen, pong.Setup)
		}
	}

	g.drawScreen(screen)
	return nil
}

func (g *Game) restart(screen *ebiten.Image, state pong.GameState) {
	screenWidth, _ := screen.Size()
	g.state = state
	g.speed = 0

	if state == pong.Setup {
		g.player.Score = 0
		g.ai.Score = 0
	}
	g.player.Pos = pong.Pos{
		X: pong.PaddleStart, Y: pong.Center(screen).Y}
	g.ai.Pos = pong.Pos{
		X: float64(screenWidth - pong.PaddleStart - pong.PaddleWidth), Y: pong.Center(screen).Y}
	g.ball.Pos = pong.Center(screen)
	g.ball.VelX = pong.BallVelX
	g.ball.VelY = pong.BallVelY
}

func main() {
	game := openTheGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

}
