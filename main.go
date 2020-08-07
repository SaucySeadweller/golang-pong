package main

import (
	"github.com/SaucySeadweller/golang-pong/pong"

	"github.com/hajimehoshi/ebiten"
)

type game struct {
	Game *pong.Game
}

const (
	vel         = 5.0
	paddleSpeed = 10
	update      = 15
	accel       = 0.5
)

func openTheGame() *game {
	game := &game{Game: &pong.Game{}}
	return game
}

func (g *game) init(screen *ebiten.Image) {
	g.Game.InitializeGame(screen)
}

func main() {

	game := openTheGame()
	if err := ebiten.RunGame(game.Game); err != nil {
		panic(err)
	}

}
