package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	im := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for y := im.Bounds().Min.Y; y < im.Bounds().Max.Y; y++ {
		for x := im.Bounds().Min.X; x < im.Bounds().Max.X; x++ {
			im.SetRGBA(x, y, color.RGBA{
				uint8(y),
				uint8(y * x % 256),
				uint8(x),
				255})
		}
	}
	if file, err := os.Create("im.png"); err == nil {
		if err := png.Encode(file, im); err == nil {
			fmt.Println("success")
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
