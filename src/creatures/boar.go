package creatures

type Boar struct {
	SizeScaler  float64
	PicturePath string
}

func NewBoar() *Boar {
	newBoar := Boar{
		SizeScaler:  .35,
		PicturePath: "../images/boar.png",
	}
	return &newBoar
}
