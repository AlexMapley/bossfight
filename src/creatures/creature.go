package creatures

import "github.com/faiface/pixel"

// Base Image for all creatures
type Creature struct {
	PicturePath string
	Position    pixel.Vec
	SizeScaler  float64
	Health      int
}
