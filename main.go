package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"log"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

// func loadImage(paths []string) <-chan Job {
// 	out := make(chan Job)
// 	go func() {
// 		// For each input path create a job and add it to
// 		// the out channel
// 		for _, p := range paths {
// 			job := Job{InputPath: p,
// 				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
// 			job.Image = imageprocessing.ReadImage(p)
// 			out <- job
// 		}
// 		close(out)
// 	}()
// 	return out
// }

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			job := Job{InputPath: p,
				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			img, err := imageprocessing.ReadImage(p)
			if err != nil {
				// Log the error and continue to the next image
				log.Printf("Error loading image at path %s: %v", p, err)
				continue
			}
			job.Image = img
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {

	imagePaths := []string{"images/IMG_0259.jpeg",
		"images/IMG_2077.jpeg",
		"images/wrongpath.jpeg",
		"images/IMG_1753.jpeg",
		"images/IMG_2768.jpeg",
	}

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
}
