package main

import (
	"image"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)



func main() {
	pixelgl.Run(run)
}



func run() {

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

	dolphPic, err := loadPicture("../images/dolph.png")
	if err != nil {
		panic(err)
	}

	dolph := pixel.NewSprite(dolphPic, dolphPic.Bounds())

	win.Clear(colornames.Skyblue)

	// Pixe.IM stands for Pixel "Identity Matrix"
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	mat.ScaledXY(win.Bounds().Center(), pixel.V(0.5, 0.5))
	dolph.Draw(win, mat)

	
	// Sprite attributes
	dolphAngle := 0.0

	// Global fields
	lastFrameTime := time.Now()

	// Main window loop
  for !win.Closed() {

		// Global fields
		deltatTime := time.Since(lastFrameTime).Seconds()
		lastFrameTime = time.Now()

		// Clear screen
		win.Clear(colornames.Skyblue)

		// Attribute updates
		dolphAngle += 1.5 * deltatTime



		// Drawing Dolph
		mat := pixel.IM
		mat = mat.Rotated(pixel.ZV, dolphAngle)
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
