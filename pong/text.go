package pong

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
)

const (
	fontSize  = 30
	smallSize = int(fontSize / 2)
)

var (
	//Font is the Normal font
	Font font.Face
	//SmallFont is the small font
	SmallFont font.Face
)

//InitFont initializes the fonts
func InitFont() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	var dpi float64 = 72
	Font = truetype.NewFace(tt, &truetype.Options{
		Size:    float64(fontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	SmallFont = truetype.NewFace(tt, &truetype.Options{
		Size:    float64(smallSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

//BigTextDraw draws the big text
func BigTextDraw(state GameState, color color.Color, screen *ebiten.Image) {

	screenWidth, _ := screen.Size()
	var textT []string

	switch state {
	case Key:
		textT = []string{
			"",
			"Player",
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
	case GameOver:
		textT = []string{
			"",
			"Game Over",
			"Esc = quit",
			"Space = play again",
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
	if state == Play || state == Key || state == Setup {
		message = append(message, "PONG")
	}
	for i, length := range message {
		s := (screenWidth - len(length)*smallSize) / 2

		text.Draw(screen, length, Font, s, screenHeight-500+(i-2)*fontSize, color)

	}
}

//Center finds the center of the screen
func Center(screen *ebiten.Image) Pos {
	screenHeight, screenWidth := screen.Size()
	return Pos{
		X: float64(screenWidth / 2),
		Y: float64(screenHeight / 2),
	}
}
