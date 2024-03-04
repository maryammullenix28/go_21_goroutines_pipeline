package main_test

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"os"
	"testing"
)

func TestReadImage(t *testing.T) {
	path := "images/IMG_0259.jpeg"

	img, _ := imageprocessing.ReadImage(path)
	if img == nil {
		t.Error("ReadImage returned nil")
	}
}

func TestWriteImage(t *testing.T) {
	path := "images/IMG_0259.jpeg"
	img, _ := imageprocessing.ReadImage(path)
	newPath := "test_images/test_output.jpg"

	// Write the image
	imageprocessing.WriteImage(newPath, img)

	// Check if the file exists
	_, err := os.Stat(newPath)
	if os.IsNotExist(err) {
		t.Error("WriteImage did not create file")
	}

	// Clean up by deleting the created file
	err = os.Remove(newPath)
	if err != nil {
		t.Errorf("Error removing test output file: %v", err)
	}
}

func TestResize(t *testing.T) {
	// Load a test image
	img, _ := imageprocessing.ReadImage("images/IMG_0259.jpeg")
	if img == nil {
		t.Error("Error loading test image")
	}

	// Resize the image
	resizedImg := imageprocessing.Resize(img)

	// Ensure the new dimensions match the desired size
	if resizedImg.Bounds().Max.X != 500 || resizedImg.Bounds().Max.Y != 500 {
		t.Errorf("Resized image has wrong dimensions: %d x %d", resizedImg.Bounds().Max.X, resizedImg.Bounds().Max.Y)
	}
}
