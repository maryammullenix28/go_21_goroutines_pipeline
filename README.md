# Data Pipelines with Concurrency
For this assignment, students have been tasked with replicating a data pipeline example to help in exploring the possibilities of concurrency. The assignment consisted of the following steps:
- Clone the GitHub repository for image processing. 
- Build and run the program in its original form.
- Add error checking for image file input and output.
- Replace the four input image files with files of your choosing.
- Add unit tests to the code repository.
- Add benchmark methods for capturing pipeline throughput times.
- Make additional code modifications as you see fit.
- Build, test, and run the pipeline program.
- Prepare a complete README.md file documenting your work.

The program we're looking at performs image manipulations, specifically makes them grayscale and resizes to 500x500 pixels.
### Error Checking
When implementing error checking for image file input and outputs, I chose to add statements that specified what error occurred (error with opening, decoding, creating, or encoding) as well as a file type checker. This makes sure to check for valid image filetypes (JPG, JPEG, PNG). These changes were made in the image_processing.go file.
### Unit Testing
The unit test written assess the functionality of three key functions in the program - readImage(), writeImage(), and resize(). 
- `TestReadImage` ensures that the ReadImage function correctly reads an image file from the provided path.
- `TestWriteImage` assesses that the WriteImage function correctly writes an image file to the specified path.
- `TestResize` looks at the Resize function so that it correctly resizes an image to 500x500 pixels.
### Benchmark Methods

### Additional Code Modifications
