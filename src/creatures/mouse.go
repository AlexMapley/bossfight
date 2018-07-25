package creatures

type Mouse struct {
	Base        *Creature
	SizeScaler  float64
	PicturePath string
	Health      int
}

func NewMouse() *Mouse {
	newMouse := Mouse{
		SizeScaler:  .15,
		PicturePath: "../images/mouse.png",
		Health:      5,
	}
	return &newMouse
}
