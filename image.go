package main

import (
	"image"
	"image/jpeg"
	// "image/png"
	"os"
)

func main() {
	// Open image file
	file, err := os.Open("input.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// Create output file
	out, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Encode as JPEG
	jpeg.Encode(out, img, &jpeg.Options{Quality: 90})
}
