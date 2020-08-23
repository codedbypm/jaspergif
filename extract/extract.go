package extract

import (
	"log"
	"gocv.io/x/gocv"
 )

 // StartExtraction is the new awesome thing
func StartExtraction() error {

	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
    	log.Fatalf("error opening device: %v", err)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	window := gocv.NewWindow("webcamwindow")
	defer window.Close()

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from the webcam")
			continue
		}

		window.IMShow(img)
		window.WaitKey(50)
	return nil
}
