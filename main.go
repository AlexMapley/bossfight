package main

import (
	"image"
	"os"

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

	dolphPic, err := loadPicture("images/dolph.png")
	if err != nil {
		panic(err)
	}

	dolph := pixel.NewSprite(dolphPic, dolphPic.Bounds())

	win.Clear(colornames.Skyblue)

	// Pixe.IM stands for Pixel "Identity Matrix"
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	dolph.Draw(win, mat)

	
	// Sprite attributes
	dolphAngle := 0.0

	// Main window loop
  for !win.Closed() {

		// Clear screen
		win.Clear(colornames.Skyblue)

		// Attribute updates
		dolphAngle += 0.05



		mat := pixel.IM
		mat = mat.Rotated(pixel.ZV, dolphAngle)
		mat = mat.Moved(win.Bounds().Center())

		dolph.Draw(win, mat)
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
