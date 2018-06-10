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

	mainCharacter, err := loadPicture("images/dolph.png")
	if err != nil {
		panic(err)
	}

	dolph := pixel.NewSprite(mainCharacter, mainCharacter.Bounds())

	win.Clear(colornames.Skyblue)
	dolph.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

  for !win.Closed() {
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
