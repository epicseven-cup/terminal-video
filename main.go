package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"gocv.io/x/gocv"
	"os"
	"video-stream-terminal/pkg"
)

func main() {

	deviceID := 0

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Println("Error opening video capture device ", err)
	}
	// Closing the webcam
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	go func() {
		for {
			if ok := webcam.Read(&img); !ok {
				fmt.Println("Device closed")
				return
			}

			//gocv.Demosaicing(img, &grayImag, 7)
			gocv.IMWrite("gray.jpg", img)
		}
	}()

	p := tea.NewProgram(pkg.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
