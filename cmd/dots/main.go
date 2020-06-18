package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/glibsm/dots"
)

func main() {
	var (
		w int
		s string
		i bool
	)

	flag.Usage = func() {
		fmt.Println("Usage:", os.Args[0], "-s <FILEPATH> <flags>")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	flag.IntVar(&w, "w", 100, "Width (in characters)")
	flag.BoolVar(&i, "i", false, "Invert the image")
	flag.StringVar(&s, "s", "", "Image source")
	flag.Parse()

	if s == "" {
		log.Println("-s flag is required")
		flag.Usage()
		os.Exit(1)
	}

	bytes, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes)
	if err != nil {
		log.Fatal(err)
	}

	var opts []dots.Option
	if i {
		opts = append(opts, dots.Invert())
	}
	opts = append(opts, dots.Width(w))

	if err := dots.Write(img, opts...); err != nil {
		log.Fatal(err)
	}
}
