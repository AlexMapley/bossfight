package main

import (
	"image"
	"os"
	"fmt"
	"math"

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

	mainCharacter, err := loadPicture("images/dolph.png")
	if err != nil {
		panic(err)
	}

	dolph := pixel.NewSprite(mainCharacter, mainCharacter.Bounds())

	win.Clear(colornames.Skyblue)

	// Pixe.IM stands for Pixel "Identity Matrix"
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	dolph.Draw(win, mat)

	u := pixel.V(2.7, 5)
	v := pixel.V(10, 3.14)
	w := u.Add(v)
	fmt.Println(w.X) // 12.7

  for !win.Closed() {
		mat = mat.Rotated(win.Bounds().Center(), math.Pi/64)
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
