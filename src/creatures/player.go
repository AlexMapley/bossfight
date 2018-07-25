package creatures

import "github.com/faiface/pixel"

type Player struct {
	Base        *Creature
	PicturePath string
	Position    pixel.Vec
	SizeScaler  float64
	Health      int
}

func NewPlayer() *Player {
	newPlayer := Player{
		SizeScaler:  .15,
		PicturePath: "../images/dolph.png",
		Health:      10,
	}
	return &newPlayer
}
