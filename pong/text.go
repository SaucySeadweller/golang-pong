package pong

import (
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"

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

//Center finds the center of the screen
func Center(screen *ebiten.Image) Pos {
	screenHeight, screenWidth := screen.Size()
	return Pos{
		X: float64(screenWidth / 2),
		Y: float64(screenHeight / 2),
	}
}
