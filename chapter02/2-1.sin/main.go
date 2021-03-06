package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	const size = 300

	pic := image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		for y := 0; y < size; y++ {
			y := size/2 - math.Sin(s)*size/2
			pic.SetGray(x, int(y), color.Gray{255})
		}
	}

	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, pic)
	file.Close()
}
