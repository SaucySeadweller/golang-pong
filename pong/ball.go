package pong

type Ball struct {
	size   int
	accelX int
	accelY int
	velX   float64
	velY   float64
	posX   float64
	posY   float64
}

const (
	BallSize   = 10
	BallAccelX = 1
	BallAccelY = 1
	BallPosX   = 50
	BallPosY   = 50
	BallVelX   = 2
	BallVelY   = 2
)
