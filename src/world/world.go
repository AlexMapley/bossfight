package world

import (
	"image"
	"math/rand"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Global environment variables
var (
	Trees    		[]*pixel.Sprite
	Enemies 		[]*pixel.Sprite
	Matrices 		[]pixel.Matrix
	EnemyMatrices	[]pixel.Matrix
)



func LoadTrees(win pixelgl.Window) {
	// Load in trees
	treeSheet, err := loadPicture("../images/trees.png")
	if err != nil {
		panic(err)
	}

	// Selects tree images from pixel sheet
	var treesFrames []pixel.Rect
	for x := treeSheet.Bounds().Min.X; x < treeSheet.Bounds().Max.X; x += 32 {
		for y := treeSheet.Bounds().Min.Y; y < treeSheet.Bounds().Max.Y; y += 32 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
		}
	}

	// Random Tree generator
	for i := 0; i < 1000; i++ {

		tree := pixel.NewSprite(treeSheet, treesFrames[rand.Intn(len(treesFrames))])
		Trees = append(Trees, tree)

		xValue := rand.Intn(2000)
		yValue := rand.Intn(2000)
		xInversion := rand.Intn(2)
		yInversion := rand.Intn(2)
		sizeScaler := float64(rand.Intn(10)) + 1
		if (xInversion <= 1) {
			xValue *= -1 
		}
		if (yInversion <= 1) {
			yValue *= -1 
		}
		placementVector := pixel.V(float64(xValue), float64(yValue))
		Matrices = append(Matrices, pixel.IM.Scaled(pixel.ZV, sizeScaler).Moved((win.Bounds().Center().Add(placementVector)).Scaled(5)	))
	}
	return
}


func EnemyGenerator(win pixelgl.Window) {
	// Enemy Generator
	enemyPic, err := loadPicture("../images/alex.png")
	if err != nil {
		panic(err)
	}
	iteratorNumber2 := 0
	for (iteratorNumber2 <= 10) {
		iteratorNumber2 += 1
		
		enemy := pixel.NewSprite(enemyPic, enemyPic.Bounds())
		Enemies = append(Enemies,enemy)
		xValue := rand.Intn(2000)
		yValue := rand.Intn(2000)
		xInversion := rand.Intn(2)
		yInversion := rand.Intn(2)
		//sizeScaler := float64(rand.Intn(2)) + .2
		if (xInversion <= 1) {
			xValue *= -1 
		}
		if (yInversion <= 1) {
			yValue *= -1 
		}
		placementVector := pixel.V(float64(xValue), float64(yValue))
		EnemyMatrices = append(EnemyMatrices, pixel.IM.Scaled(pixel.ZV, 1).Moved((win.Bounds().Center().Add(placementVector)).Scaled(5)	))
		Matrices = append(Matrices, pixel.IM.Scaled(pixel.ZV, 1).Moved((win.Bounds().Center().Add(placementVector)).Scaled(5)	))
	}
	
	return
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
