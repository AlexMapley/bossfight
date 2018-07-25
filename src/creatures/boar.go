package creatures

type Boar struct {
	Base        *Creature
	SizeScaler  float64
	PicturePath string
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
