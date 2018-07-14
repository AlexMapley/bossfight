package main

import (
	"image"
	"os"
	"time"
	"math"
	"fmt"

	_ "image/png"

	"./world"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	//"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	//"golang.org/x/image/font/basicfont"
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
	world.LoadTrees(*win)
	world.EnemyGenerator(*win)

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
		for i, tree := range world.Trees {
			tree.Draw(win, world.Matrices[i])
		}

		// Plant Enemies
		for i, enemy := range world.Enemies {
			enemy.Draw(win, world.Matrices[i+len(world.Trees)])
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
			for i := range world.Enemies {

				// TODO: Add absolute value function here
				if (world.Matrices[i][4] - playerLocation.X < 15 && world.Matrices[i][5] - playerLocation.Y < 15 && world.Matrices[i][4] - playerLocation.X > -15 && world.Matrices[i][5] - playerLocation.Y > -15) {
					if (world.Matrices[i][4] - playerLocation.X > 0) {
						fmt.Println("Case1")
						fmt.Println("world.Matrices[i][4] ",world.Matrices[i][4] )
						fmt.Println("playerLocation.X ", playerLocation.X)
						fmt.Println("world.Matrices[i][5] ",world.Matrices[i][5] )
						fmt.Println("playerLocation.Y ", playerLocation.Y)
						collision = true
					}
				}
			}

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
			for i := range world.Enemies {
				
				// TODO: Add absolute value function here
				if (world.Matrices[i][4] - playerLocation.X < 15 && world.Matrices[i][5] - playerLocation.Y < 15 && world.Matrices[i][4] - playerLocation.X > -15 && world.Matrices[i][5] - playerLocation.Y > -15) {
					if (world.Matrices[i][4] - playerLocation.X < 0) {
						fmt.Println("Case2")
						fmt.Println("world.Matrices[i][4] ",world.Matrices[i][4] )
						fmt.Println("playerLocation.X ", playerLocation.X)
						fmt.Println("world.Matrices[i][5] ",world.Matrices[i][5] )
						fmt.Println("playerLocation.Y ", playerLocation.Y)
						collision = true
					}
				}
			}
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
			for i := range world.Enemies {
				
				// TODO: Add absolute value function here
				if (world.Matrices[i][4] - playerLocation.X < 15 && world.Matrices[i][5] - playerLocation.Y < 15 && world.Matrices[i][4] - playerLocation.X > -15 && world.Matrices[i][5] - playerLocation.Y > -15) {
					if (world.Matrices[i][5] - playerLocation.Y > 0) {
						fmt.Println("Case3")
						fmt.Println("world.Matrices[i][4] ",world.Matrices[i][4] )
						fmt.Println("playerLocation.X ", playerLocation.X)
						fmt.Println("world.Matrices[i][5] ",world.Matrices[i][5] )
						fmt.Println("playerLocation.Y ", playerLocation.Y)
						collision = true
					}
				}
			}
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
			for i := range world.Enemies {
				
				// TODO: Add absolute value function here
				if (world.Matrices[i][4] - playerLocation.X < 15 && world.Matrices[i][5] - playerLocation.Y < 15 && world.Matrices[i][4] - playerLocation.X > -15 && world.Matrices[i][5] - playerLocation.Y > -15) {
					if (world.Matrices[i][5] - playerLocation.Y < 0) {
						fmt.Println("Case4")
						fmt.Println("world.Matrices[i][4] ",world.Matrices[i][4] )
						fmt.Println("playerLocation.X ", playerLocation.X)
						fmt.Println("world.Matrices[i][5] ",world.Matrices[i][5] )
						fmt.Println("playerLocation.Y ", playerLocation.Y)
						collision = true
					}
				}
			}
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

