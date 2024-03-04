package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func ReadImage(path string) (image.Image, error) {

	inputFile, err := os.Open(path)
	if err != nil {
		fmt.Errorf("Error opening file: %v", err)
	}
	defer inputFile.Close()

	// Check file extension
	fileExt := strings.ToLower(filepath.Ext(path))
	if fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".png" {
		return nil, fmt.Errorf("invalid file type, must be .jpg, .jpeg, or .png")
	}

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Errorf("Error decoding file: %v", err)
		fmt.Println(path)
	}
	return img, err
}

func WriteImage(path string, img image.Image) {
	outputFile, err := os.Create(path)
	if err != nil {
		fmt.Errorf("Error creating file: %v", err)
		panic(err)
	}
	defer outputFile.Close()

	// Encode the image to the new file
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		fmt.Errorf("Error encoding file: %v", err)
		panic(err)
	}
}

func Grayscale(img image.Image) image.Image {
	// Create a new grayscale image
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

func Resize(img image.Image) image.Image {
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return resizedImg
}
