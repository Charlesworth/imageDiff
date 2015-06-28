package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"time"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images.
	//_ "image/jpeg"
)

func main() {

	startTime := time.Now()
	defer runtime(startTime)

	// img := loadJPEG("img2.jpeg")
	img := loadJPEG("testr.jpg")
	imgOld := loadJPEG("testr2.jpg")
	//img := loadJPEG("2.jpg")
	//imgOld := loadJPEG("1.jpg")

	rmvGreen := rmvGreenAndCommon(img, imgOld)

	saveImage(rmvGreen, "alchemy")

}

func runtime(startTime time.Time) {
	endTime := time.Now()
	fmt.Println("time: ", endTime.Sub(startTime))
}

func saveImage(img image.Image, fileName string) {
	finalFile, _ := os.Create(fileName + ".jpeg")
	defer finalFile.Close()
	jpeg.Encode(finalFile, img, &jpeg.Options{jpeg.DefaultQuality})
}

func rmvGreenAndCommon(img image.Image, imgOld image.Image) image.Image {
	bounds := img.Bounds()

	rmvGreen := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rmvGreen, bounds, img, bounds.Min, draw.Src)

	greens := 0
	match := 0
	white := color.RGBA{255, 255, 255, 255}

	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			r, g, b, _ := img.At(x, y).RGBA()
			//fmt.Print("x:", x, " y:", y, " r:", r, " g:", g, " b:", b) //, " lu:", luminance(r, g, b))

			//count green pixels
			if isGreen(r, g, b) {
				greens++
				rmvGreen.Set(x, y, white)

			} else {

				rOld, gOld, bOld, _ := imgOld.At(x, y).RGBA()
				if isSimilar(r, g, b, rOld, gOld, bOld) {
					match++
					rmvGreen.Set(x, y, white)

				}
			}
		}
	}

	tot := bounds.Max.Y * bounds.Max.X
	mismatch := tot - greens - match
	fmt.Println("total pixels", tot)
	fmt.Println("green pixels", greens)
	fmt.Println("matched pixels", match)
	fmt.Println("new pixels", mismatch)

	if mismatch > (tot / 10) {
		fmt.Println("**ALERT** Over 10 percent change **ALERT**")
	} else {
		fmt.Println("Change under 10 percent")
	}

	return rmvGreen
}

func isSimilar(r uint32, g uint32, b uint32, rOld uint32, gOld uint32, bOld uint32) bool {
	if (r < rOld+10000 && r > rOld-10000) && (b < bOld+10000 && b > bOld-10000) && (g < gOld+10000 && g > gOld-10000) {
		return true
	}
	return false
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
