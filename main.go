package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/mattn/go-mjpeg"
)

var url = flag.String("url", "", "Camera host")

func main() {
	flag.Parse()
	if *url == "" {
		flag.Usage()
		os.Exit(1)
	}

	dec, err := mjpeg.NewDecoderFromURL(*url)
	if err != nil {
		log.Fatal(err)
	}

	var img image.Image

	var tmp image.Image
	tmp, err = dec.Decode()
	if err != nil {
		fmt.Println(err)
	}
	img = tmp
	frame, _ := os.Create("frame.jpg")
	jpeg.Encode(frame, img, &jpeg.Options{jpeg.DefaultQuality})
}
