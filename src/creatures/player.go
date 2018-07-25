package creatures

type Player struct {
	SizeScaler  float64
	PicturePath string
}

func NewPlayer() *Player {
	newPlayer := Player{
		SizeScaler:  .15,
		PicturePath: "../images/dolph.png",
	}
	return &newPlayer
}
