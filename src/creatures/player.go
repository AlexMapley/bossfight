package creatures

type Player struct {
	Base        *Creature
	SizeScaler  float64
	PicturePath string
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
