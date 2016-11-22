package main

import (
	"flag"
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

	img, err := dec.Decode()
	if err != nil {
		log.Fatal(err)
	}
	frame, err := os.Create("frame.jpg")
	if err != nil {
		log.Fatal(err)
	}

	jpeg.Encode(frame, img, &jpeg.Options{jpeg.DefaultQuality})
}
