package main

import (
	"image"
	"os"
	"time"
	"math"
	"math/rand"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	//"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	//"golang.org/x/image/font/basicfont"
)



func main() {
	pixelgl.Run(run)
}



func run() {

  cfg := pixelgl.WindowConfig{
		Title:  "Boss Fight!",
		Bounds: pixel.R(0, 0, 1024, 1024),
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
	var (
		trees    []*pixel.Sprite
		matrices []pixel.Matrix
	)

	// Load in main character
	dolphPic, err := loadPicture("../images/dolph.png")
	if err != nil {
		panic(err)
	}
	dolph := pixel.NewSprite(dolphPic, dolphPic.Bounds())


	// Pixe.IM stands for Pixel "Identity Matrix"
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	mat.ScaledXY(win.Bounds().Center(), pixel.V(0.5, 0.5))
	dolph.Draw(win, mat)

	
	// Camera field
	var (
		camPos   = pixel.ZV
		camSpeed = 500.0
		camZoom      = 1.0
		camZoomSpeed = 1.2
	)
	
	// Sprite attributes
	dolphAngle := 0.0

	// Global fields
	lastFrameTime := time.Now()

	// Main window loop
  	for !win.Closed() {

		// Clear screen
		win.Clear(colornames.Whitesmoke)

		// Global fields
		deltatTime := time.Since(lastFrameTime).Seconds()
		lastFrameTime = time.Now()

		// Plant trees
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			tree := pixel.NewSprite(treeSheet, treesFrames[rand.Intn(len(treesFrames))])
			trees = append(trees, tree)
			matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 2).Moved(win.MousePosition()))
		}
		for i, tree := range trees {
			tree.Draw(win, matrices[i])
		}

		// Moving the character
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

		

		// Attribute updates
		dolphAngle += 1.5 * deltatTime



		// Drawing Dolph
		mat := pixel.IM
		//mat = mat.Rotated(pixel.ZV, dolphAngle)
		mat = mat.Moved(win.Bounds().Center())
		mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(0.25, 0.25))
		dolph.Draw(win, mat)


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
