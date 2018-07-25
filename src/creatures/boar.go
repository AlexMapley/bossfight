package creatures

import "github.com/faiface/pixel"

type Boar struct {
	Base        *Creature
	PicturePath string
	Position    pixel.Vec
	SizeScaler  float64
	Health      int
}

func NewBoar() *Boar {
	newBoar := Boar{
		SizeScaler:  .35,
		PicturePath: "../images/boar.png",
		Health:      25,
	}
	return &newBoar
}
