package main

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images.
	_ "image/jpeg"
)

func main() {

	img := loadJPEG("img2.jpeg")
	bounds := img.Bounds()

	rmvGreen := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rmvGreen, bounds, img, bounds.Min, draw.Src)

	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.

	greens := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			r, g, b, _ := img.At(x, y).RGBA()
			fmt.Print("x:", x, " y:", y, " r:", r, " g:", g, " b:", b, " lu:", luminance(r, g, b))

			//count green pixels
			if isGreen(r, g, b) {
				greens++
				fmt.Println("    GREEN!!!!!")
				//rmvGreen.Set(x, y, color here)
			} else {
				fmt.Println()
			}

		}
	}
	fmt.Println("greens", greens)
}

//isGreen returns a bool if the RBG input equates to a green color
func isGreen(r uint32, g uint32, b uint32) bool {
	if g > r && g > b && g > 10000 {
		return true
	}
	return false
}

//credit code from https://github.com/esdrasbeleza/blzimg
//luminance returns the luminance out ~65021 for RGB values r, g, b
func luminance(r uint32, g uint32, b uint32) uint32 {
	return uint32(0.2126*float32(r) + 0.7152*float32(g) + 0.0722*float32(b))
}

//loadJPEG decodes the .jpeg file "fileName" and returns an image.Image
func loadJPEG(fileName string) image.Image {
	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	return img
}
