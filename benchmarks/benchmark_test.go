package image_processing

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"testing"
)

func BenchmarkReadImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Read the image
		_, _ = imageprocessing.ReadImage("test_images/1348996.jpg")
	}
}

func BenchmarkWriteImage(b *testing.B) {
	img, _ := imageprocessing.ReadImage("test_images/1348996.jpg")
	for i := 0; i < b.N; i++ {
		// Write the image
		imageprocessing.WriteImage("test_images/test_output.jpg", img)
	}
}

func BenchmarkGrayscale(b *testing.B) {
	img, _ := imageprocessing.ReadImage("test_images/1348996.jpg")
	for i := 0; i < b.N; i++ {
		// Convert to grayscale
		_ = imageprocessing.Grayscale(img)
	}
}

func BenchmarkResize(b *testing.B) {
	img, _ := imageprocessing.ReadImage("test_images/1348996.jpg")
	for i := 0; i < b.N; i++ {
		// Resize the image
		_ = imageprocessing.Resize(img)
	}
}
