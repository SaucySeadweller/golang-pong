package pong

type Paddle struct {
	width  int
	length int
	posX   float64
	posY   float64
	velX   float64
	velY   float64
}

const (
	PaddleWidth  = 10
	PaddleLength = 30
	PaddleVelY   = 2
	PaddleVelX   = 2
	PaddlePosX   = 0
	PaddlePosY   = 15
)
