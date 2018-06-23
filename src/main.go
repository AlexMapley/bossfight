package main

import (
	"image"
	"os"
	"time"
	"math"
	"math/rand"
	"fmt"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	//"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	//"golang.org/x/image/font/basicfont"
)

// Global environment variables
var (
	trees    []*pixel.Sprite
	matrices []pixel.Matrix
)

// Main calls the program loop
func main() {
	pixelgl.Run(run)
}



// Actual executable code
func run() {


  fmt.Println("game started")


  cfg := pixelgl.WindowConfig{
		Title:  "Boss Fight!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

  win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Removes pixelation effect on sprite movement
	win.SetSmooth(true)


	// Load in trees
	treeSheet, err := loadPicture("../images/trees.png")
	if err != nil {
		panic(err)
	}
	var treesFrames []pixel.Rect
	for x := treeSheet.Bounds().Min.X; x < treeSheet.Bounds().Max.X; x += 32 {
		for y := treeSheet.Bounds().Min.Y; y < treeSheet.Bounds().Max.Y; y += 32 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
		}
	}

	// Load in main character
	dolphPic, err := loadPicture("../images/dolph.png")
	if err != nil {
		panic(err)
	}
	dolph := pixel.NewSprite(dolphPic, dolphPic.Bounds())


	// Player Position Attributes
	playerLocation := win.Bounds().Center()
	playerAngle := 0.0
	
	// Camera field
	var (
		camPos 		= win.Bounds().Center()
		camSpeed 	= 300.0
		camZoom      = 1.0
		camZoomSpeed = 1.05
	)

	// Background elements
		// Random Tree generator
		iteratorNumber := 0
		for (iteratorNumber <= 1000) {
			iteratorNumber += 1

			tree := pixel.NewSprite(treeSheet, treesFrames[rand.Intn(len(treesFrames))])
			trees = append(trees, tree)

			xValue := rand.Intn(2000)
			yValue := rand.Intn(2000)
			xInversion := rand.Intn(2)
			yInversion := rand.Intn(2)
			if (xInversion <= 1) {
				xValue *= -1 
			}
			if (yInversion <= 1) {
				yValue *= -1 
			}
			placementVector := pixel.V(float64(xValue), float64(yValue))
			matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 2).Moved(	(win.Bounds().Center().Add(placementVector)).Scaled(5)	))
		}
	

	// Global fields
	lastFrameTime := time.Now()

	// Main window loop
  	for !win.Closed() {

		// Clear screen
		win.Clear(colornames.Skyblue)

		// Global fields
		deltatTime := time.Since(lastFrameTime).Seconds()
		lastFrameTime = time.Now()

		
		
		// Plant trees
		for i, tree := range trees {
			tree.Draw(win, matrices[i])
		}

		// Moving the camera
		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)
		if win.Pressed(pixelgl.KeyLeft) {
			camPos.X -= camSpeed * deltatTime
		}
		if win.Pressed(pixelgl.KeyRight) {
			camPos.X += camSpeed * deltatTime
		}
		if win.Pressed(pixelgl.KeyDown) {
			camPos.Y -= camSpeed * deltatTime
		}
		if win.Pressed(pixelgl.KeyUp) {
			camPos.Y += camSpeed * deltatTime
		}
		if win.Pressed(pixelgl.KeyUp) {
			camPos.Y += camSpeed * deltatTime
		}
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)




		// Moving the Character

		// Strafing Left
		if win.Pressed(pixelgl.KeyA) {
			// Prevents tree collision
			collision := false
			// for i := range trees {

			// 	// TODO: Add absolute value function here
			// 	if (matrices[i][4] - playerLocation.X < 15 && matrices[i][5] - playerLocation.Y < 15 && matrices[i][4] - playerLocation.X > -15 && matrices[i][5] - playerLocation.Y > -15) {
			// 		if (matrices[i][4] - playerLocation.X > 0) {
			// 			fmt.Println("Case1")
			// 			fmt.Println("Matrices[i][4] ",matrices[i][4] )
			// 			fmt.Println("playerLocation.X ", playerLocation.X)
			// 			fmt.Println("Matrices[i][5] ",matrices[i][5] )
			// 			fmt.Println("playerLocation.Y ", playerLocation.Y)
			// 			collision = true
			// 		}
			// 	}
			// }

			if (!collision) {
				playerLocation.X -= 5
			}

			// Camera follows
			camPos.X -= camSpeed * deltatTime
		}

		// Strafing Right
		if win.Pressed(pixelgl.KeyD) {
			// Prevents tree collision
			collision := false
			// for i := range trees {
				
			// 	// TODO: Add absolute value function here
			// 	if (matrices[i][4] - playerLocation.X < 15 && matrices[i][5] - playerLocation.Y < 15 && matrices[i][4] - playerLocation.X > -15 && matrices[i][5] - playerLocation.Y > -15) {
			// 		if (matrices[i][4] - playerLocation.X < 0) {
			// 			fmt.Println("Case2")
			// 			fmt.Println("Matrices[i][4] ",matrices[i][4] )
			// 			fmt.Println("playerLocation.X ", playerLocation.X)
			// 			fmt.Println("Matrices[i][5] ",matrices[i][5] )
			// 			fmt.Println("playerLocation.Y ", playerLocation.Y)
			// 			collision = true
			// 		}
			// 	}
			// }
			if (!collision) {
				playerLocation.X += 5
			}

			// Camera follows
			camPos.X += camSpeed * deltatTime
		}

		// Moving Down
		if win.Pressed(pixelgl.KeyS) {
			// Prevents tree collision
			collision := false
			// for i := range trees {
				
			// 	// TODO: Add absolute value function here
			// 	if (matrices[i][4] - playerLocation.X < 15 && matrices[i][5] - playerLocation.Y < 15 && matrices[i][4] - playerLocation.X > -15 && matrices[i][5] - playerLocation.Y > -15) {
			// 		if (matrices[i][5] - playerLocation.Y > 0) {
			// 			fmt.Println("Case3")
			// 			fmt.Println("Matrices[i][4] ",matrices[i][4] )
			// 			fmt.Println("playerLocation.X ", playerLocation.X)
			// 			fmt.Println("Matrices[i][5] ",matrices[i][5] )
			// 			fmt.Println("playerLocation.Y ", playerLocation.Y)
			// 			collision = true
			// 		}
			// 	}
			// }
			if (!collision) {
				playerLocation.Y -= 5
			}

			// Camera follows
			camPos.Y -= camSpeed * deltatTime
		}

		// Moving Up
		if win.Pressed(pixelgl.KeyW) {
			// Prevents tree collision
			collision := false
			// for i := range trees {
				
			// 	// TODO: Add absolute value function here
			// 	if (matrices[i][4] - playerLocation.X < 15 && matrices[i][5] - playerLocation.Y < 15 && matrices[i][4] - playerLocation.X > -15 && matrices[i][5] - playerLocation.Y > -15) {
			// 		if (matrices[i][5] - playerLocation.Y < 0) {
			// 			fmt.Println("Case4")
			// 			fmt.Println("Matrices[i][4] ",matrices[i][4] )
			// 			fmt.Println("playerLocation.X ", playerLocation.X)
			// 			fmt.Println("Matrices[i][5] ",matrices[i][5] )
			// 			fmt.Println("playerLocation.Y ", playerLocation.Y)
			// 			collision = true
			// 		}
			// 	}
			// }
			if (!collision) {
				playerLocation.Y += 5
			}

			// Camera follows
			camPos.Y += camSpeed * deltatTime
		}

		// Rotating Left
		if win.Pressed(pixelgl.KeyQ) {
			playerAngle -= .02
		}

		// Rotating Right
		if win.Pressed(pixelgl.KeyE) {
			playerAngle += .02
		}




		// Drawing Dolph
		playerPosition := pixel.IM
		playerPosition = playerPosition.ScaledXY(playerLocation, pixel.V(0.1, 0.1))
		playerPosition = playerPosition.Rotated(playerLocation, playerAngle)
		dolph.Draw(win, playerPosition)
		


		// Update window
		win.Update()
	}
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

