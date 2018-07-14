package creatures

type Mouse struct {
	SizeScaler  float64
	PicturePath string
}

func NewMouse() *Mouse {
	newMouse := Mouse{
		SizeScaler:  .25,
		PicturePath: "../images/mouse.png",
	}
	return &newMouse
}
