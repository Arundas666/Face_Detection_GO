// package main

// import (
// 	"fmt"
// 	"image/color"
// 	"os"

// 	"gocv.io/x/gocv"
// )

// // func main() {
// // 	// set to use a video capture device 0
// // 	deviceID := 0

// // 	// open webcam
// // 	webcam, err := gocv.OpenVideoCapture(deviceID)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	defer webcam.Close()

// // 	// open display window
// // 	window := gocv.NewWindow("Face Detect")
// // 	defer window.Close()

// // 	// prepare image matrix
// // 	img := gocv.NewMat()
// // 	defer img.Close()

// // 	// color for the rect when faces detected
// // 	blue := color.RGBA{0, 0, 255, 0}

// // 	// load classifier to recognize faces
// // 	classifier := gocv.NewCascadeClassifier()
// // 	defer classifier.Close()

// // 	if !classifier.Load("data/haarcascade_frontalface_default.xml") {
// // 		fmt.Println("Error reading cascade file: data/haarcascade_frontalface_default.xml")
// // 		return
// // 	}

// // 	fmt.Printf("start reading camera device: %v\n", deviceID)
// // 	for {
// // 		if ok := webcam.Read(&img); !ok {
// // 			fmt.Printf("cannot read device %v\n", deviceID)
// // 			return
// // 		}
// // 		if img.Empty() {
// // 			continue
// // 		}

// // 		// detect faces
// // 		rects := classifier.DetectMultiScale(img)
// // 		fmt.Printf("found %d faces\n", len(rects))

// // 		// draw a rectangle around each face on the original image
// // 		for _, r := range rects {
// // 			gocv.Rectangle(&img, r, blue, 3)
// // 		}
// // 		go ImageCapturing()
// // 		go matching()
// // 		// show the image in the window, and wait 1 millisecond
// // 		window.IMShow(img)
// // 		window.WaitKey(1)
// // 	}
// // }

// func ImageCapturing() {
// 	if len(os.Args) < 3 {
// 		fmt.Println("How to run:\n\tsaveimage [camera ID] [image file]")
// 		return
// 	}

// 	deviceID := 0
// 	saveFile := os.Args[1]

// 	webcam, err := gocv.OpenVideoCapture(deviceID)
// 	if err != nil {
// 		fmt.Printf("Error opening video capture device: %v\n", deviceID)
// 		return
// 	}
// 	defer webcam.Close()

// 	img := gocv.NewMat()
// 	defer img.Close()

// 	if ok := webcam.Read(&img); !ok {
// 		fmt.Printf("cannot read device %v\n", deviceID)
// 		return
// 	}
// 	if img.Empty() {
// 		fmt.Printf("no image on device %v\n", deviceID)
// 		return
// 	}

// 	gocv.IMWrite(saveFile, img)
// }

// func main() {

// 	if len(os.Args) != 3 {
// 		fmt.Println("Usage: feature-matching /path/to/query /path/to/train")
// 		panic("error: no files provided")
// 	}

// 	// opening query image
// 	go ImageCapturing()
// 	query := gocv.IMRead(os.Args[1], gocv.IMReadGrayScale)
// 	defer query.Close()

// 	// opening train image
// 	train := gocv.IMRead(os.Args[2], gocv.IMReadGrayScale)
// 	defer train.Close()

// 	// creating new SIFT
// 	sift := gocv.NewSIFT()
// 	defer sift.Close()

// 	// detecting and computing keypoints using SIFT method
// 	queryMask := gocv.NewMat()
// 	defer queryMask.Close()
// 	kp1, des1 := sift.DetectAndCompute(query, queryMask)
// 	defer des1.Close()

// 	trainMask := gocv.NewMat()
// 	defer trainMask.Close()
// 	kp2, des2 := sift.DetectAndCompute(train, trainMask)
// 	defer des2.Close()

// 	// finding K best matches for each descriptor
// 	bf := gocv.NewBFMatcher()
// 	matches := bf.KnnMatch(des1, des2, 2)

// 	// application of ratio test
// 	var good []gocv.DMatch
// 	for _, m := range matches {
// 		if len(m) > 1 {
// 			if m[0].Distance < 0.75*m[1].Distance {
// 				good = append(good, m[0])
// 			}
// 		}
// 	}

// 	// matches color
// 	c1 := color.RGBA{
// 		R: 0,
// 		G: 255,
// 		B: 0,
// 		A: 0,
// 	}

// 	// point color
// 	c2 := color.RGBA{
// 		R: 255,
// 		G: 0,
// 		B: 0,
// 		A: 0,
// 	}

// 	// creating empty mask
// 	mask := make([]byte, 0)

// 	// new matrix for output image
// 	out := gocv.NewMat()
// 	defer out.Close()
// 	// drawing matches
// 	gocv.DrawMatches(query, kp1, train, kp2, good, &out, c1, c2, mask, gocv.DrawDefault)

// 	// creating output window with result
// 	window := gocv.NewWindow("Output")
// 	window.IMShow(out)
// 	defer window.Close()

// 	window.WaitKey(0)
// }

package main

import (
	"fmt"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	// pre-defined training image path
	trainingImage := "./arun.jpg"

	// capture an image from the webcam
	capturedImage := captureImage()

	// compare the captured image to the training image
	matchImages(capturedImage, trainingImage)

	// clean up resources
	defer capturedImage.Close()
}

func captureImage() *gocv.Mat {
	// open webcam
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", err)
		return nil
	}
	defer webcam.Close()

	// capture an image from the webcam
	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Println("Cannot read device")
		return nil
	}
	if img.Empty() {
		fmt.Println("No image on device")
		return nil
	}
	fmt.Println("Image size:", img.Size())
	return &img
}

func matchImages(capturedImage *gocv.Mat, trainingImage string) {
	// open training image
	train := gocv.IMRead(trainingImage, gocv.IMReadGrayScale)
	defer train.Close()
	// create new SIFT
	sift := gocv.NewSIFT()
	defer sift.Close()

	// detecting and computing keypoints using SIFT method
	queryMask := gocv.NewMat()
	defer queryMask.Close()
	fmt.Println("HEyyyyy👺")
	fmt.Println(capturedImage)
	fmt.Println("HEyyyyy👺")

	kp1, des1 := sift.DetectAndCompute(*capturedImage, queryMask)
	defer des1.Close()

	trainMask := gocv.NewMat()
	defer trainMask.Close()
	kp2, des2 := sift.DetectAndCompute(train, trainMask)
	defer des2.Close()

	// finding K best matches for each descriptor
	bf := gocv.NewBFMatcher()
	matches := bf.KnnMatch(des1, des2, 2)

	// application of ratio test

	var good []gocv.DMatch
	for _, m := range matches {
		if len(m) > 1 {
			if m[0].Distance < 0.75*m[1].Distance {
				good = append(good, m[0])
			}
		}
	}

	// matches color
	c1 := color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 0,
	}

	// point color
	c2 := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 0,
	}

	// creating empty mask
	mask := make([]byte, 0)

	// new matrix for output image
	out := gocv.NewMat()
	defer out.Close()
	// drawing matches
	gocv.DrawMatches(*capturedImage, kp1, train, kp2, good, &out, c1, c2, mask, gocv.DrawDefault)

	// creating output window with result
	window := gocv.NewWindow("Output")
	window.IMShow(out)
	defer window.Close()

	window.WaitKey(0)
}
