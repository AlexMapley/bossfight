package creatures

import "github.com/faiface/pixel"

type Mouse struct {
	Base        *Creature
	PicturePath string
	Position    pixel.Vec
	SizeScaler  float64
	Health      int
}

func NewMouse() *Mouse {
	newMouse := Mouse{
		SizeScaler:  .15,
		PicturePath: "../images/mouse.png",
		Health:      5,
		Position:    pixel.V(0, 0),
	}
	return &newMouse
}
