package pong

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

func (g *Game) draw(screen *ebiten.Image) error {
	screen.Fill(Background)
	BigTextDraw(g.state, PaddleBall, screen)
	TextDraw(g.state, PaddleBall, screen)

	if g.state != Controls {
		g.player.Draw(screen, Font)
		g.ai.Draw(screen, Font)
		g.ball.Draw(screen)
	}

	return nil
}

//Update updates the screen
func (g *Game) Update(screen *ebiten.Image) error {
	g.InitializeGame(screen)
	g.StateSwitch(screen)
	g.draw(screen)

	return nil
}

//Draw draws the ball
func (b *Ball) Draw(screen *ebiten.Image) {
	ballOptions := &ebiten.DrawImageOptions{}
	ballOptions.GeoM.Translate(b.X, b.Y)
	b.Ball.Fill(PaddleBall)
	screen.DrawImage(b.Ball, ballOptions)
}

//Draw draws the paddle
func (p *Paddle) Draw(screen *ebiten.Image, f font.Face) {
	paddle := &ebiten.DrawImageOptions{}
	paddle.GeoM.Translate(p.X, p.Y-PaddleLength/2)
	p.Paddle.Fill(color.White)
	screen.DrawImage(p.Paddle, paddle)

	p.CurrentScore.x = p.X + (Center(screen).X-p.X)/2
	p.CurrentScore.y = 2 * 20

	s := strconv.Itoa(p.CurrentScore.Score)
	text.Draw(screen, s, Font, int(p.CurrentScore.x), int(p.CurrentScore.y), PaddleBall)
}

//Layout sets the window size
func (g *Game) Layout(width int, height int) (int, int) {

	return height, width
}

//BigTextDraw draws the big text
func BigTextDraw(state GameState, color color.Color, screen *ebiten.Image) {

	screenWidth, _ := screen.Size()
	var textT []string

	switch state {
	case Controls:
		textT = []string{
			"",
			"PONG",
			"",
			"W = Up",
			"S = Down",
			"",
		}
	case Setup:
		textT = []string{
			"",
			"PONG",
			"",
			"C = Controls",
			"Space = start game",
			"",
		}
	case Pause:
		textT = []string{
			"",
			"Paused",
			"",
			"Space to resume",
		}
	case GameOver:
		textT = []string{
			"",
			"Game Over",
			"Space to play again",
		}
		for i, length := range textT {
			s := (screenWidth - len(length)*fontSize) / 2
			text.Draw(screen, length, Font, s, (i+4)*fontSize, color)
		}

	}
}

//TextDraw draws little text on the screen
func TextDraw(state GameState, color color.Color, screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Size()
	message := []string{}
	if state == Play || state == Controls || state == Setup {
		message = append(message, "PONG")
	}
	for i, length := range message {
		s := (screenWidth - len(length)*smallSize) / 2

		text.Draw(screen, length, Font, s, screenHeight-500+(i-2)*fontSize, color)

	}
}
