package pong

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

//Game is a struct that defines Games attributes
type Game struct {
	state   GameState
	player  *Paddle
	ai      *Paddle
	score   int
	ball    *Ball
	speed   int
	started bool
}

const (
	vel         = 5.0
	paddleSpeed = 10
	update      = 15
	accel       = 0.5
)

const (
	//Setup preps the game state
	Setup GameState = iota
	//Play is the game state while playing
	Play
	//Pause is the gamestate while not playing
	Pause
	//GameOver is the state when the game ends
	GameOver
	//Controls is the state that shows the controls
	Controls
)

//StateSwitch cimplements the logic behind switching the states
func (g *Game) StateSwitch(screen *ebiten.Image) {
	switch g.state {
	case Setup:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = Play

		} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
			g.state = Controls
		}
	case Controls:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = Play

		}

	case Play:
		width, _ := screen.Size()

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = Pause
			break
		}
		g.player.Update(screen)
		g.ai.AI(g.ball)
		g.ai.Update(screen)

		g.ball.Update(screen, g.player, g.ai)

		if g.ball.X >= float64(width) || g.ball.X < 0 {
			g.ai.CurrentScore.Score++
			g.reset(screen, Play)
		} else if g.ball.X > float64(width) {
			g.player.CurrentScore.Score++
			g.reset(screen, Play)
		}

	case GameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.reset(screen, Setup)
		}
	}
}

func (g *Game) reset(screen *ebiten.Image, state GameState) {
	w, _ := screen.Size()
	g.state = state
	g.speed = 0

	if state == Setup {
		g.player.Score = 0
		g.ai.Score = 0
	}
	g.player.Pos = Pos{
		X: PaddleStart, Y: Center(screen).Y}
	g.ai.Pos = Pos{
		X: float64(w - PaddleStart - PaddleWidth), Y: Center(screen).Y}
	g.ball.Pos = Center(screen)
	g.ball.VelX = vel
	g.ball.VelY = vel
}

//InitializeGame sets the starting parameters for the game upon reset or startup
func (g *Game) InitializeGame(screen *ebiten.Image) {
	if g.started {
		return
	}
	g.started = true
	screenWidth, screenHeight := screen.Size()
	g.state = Setup
	g.score = 11
	g.player = &Paddle{
		Pos: Pos{
			X: PaddleStart,
			Y: float64(screenHeight / 2)},
		Score:        0,
		Speed:        paddleSpeed,
		Length:       PaddleLength,
		Width:        PaddleWidth,
		PaddleColor:  PaddleBall,
		MovementUp:   ebiten.KeyW,
		MovementDown: ebiten.KeyS,
	}
	g.ai = &Paddle{
		Pos: Pos{
			X: 469,
			Y: float64(screenHeight / 2)},
		Score:       0,
		Speed:       paddleSpeed,
		Length:      PaddleLength,
		Width:       PaddleWidth,
		PaddleColor: PaddleBall,
	}
	g.ball = &Ball{
		Pos: Pos{

			X: float64(screenWidth / 2),
			Y: float64(screenHeight / 2)},
		VelX:   vel,
		VelY:   vel,
		Radius: BallRadius,
		Color:  PaddleBall,
	}

	g.player.Paddle, _ = ebiten.NewImage(g.player.Width, g.player.Length, ebiten.FilterDefault)
	g.ai.Paddle, _ = ebiten.NewImage(g.ai.Width, g.ai.Length, ebiten.FilterDefault)
	g.ball.Ball, _ = ebiten.NewImage(int(g.ball.Radius)*2, int(g.ball.Radius)*2, ebiten.FilterDefault)
	InitFont()
}
