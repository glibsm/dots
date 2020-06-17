package dots

import (
	"fmt"
	"image"
	"io"
	"os"

	"github.com/disintegration/imaging"
	"github.com/glibsm/dots/braille"

	// import forrmats for side-effects of formats being loaded into the decoder
	_ "image/jpeg"
	_ "image/png"
)

const (
	// alpha ranges 0-0xFFFF, anything greater than half is considered visible
	halpha = 0x7FFF

	// white cut-off. For images that don't have a transparent background, any
	// shade above this value is considered to be background and therefor does
	// not get dots.
	wc = 0xEEEE
)

type writeOpts struct {
	width  int
	writer io.Writer
	invert bool
}

// Option ...
type Option func(*writeOpts) error

// Writer allows to override where the output is directed.
func Writer(w io.Writer) Option {
	return func(opts *writeOpts) error {
		opts.writer = w
		return nil
	}
}

// Width (in characters) for the length of each line of output.
func Width(w int) Option {
	return func(opts *writeOpts) error {
		opts.width = w
		return nil
	}
}

// Invert the image. Braille will become the background, and empty space will
// be used as a drawing mechanism.
func Invert() Option {
	return func(opts *writeOpts) error {
		opts.invert = true
		return nil
	}
}

// Write the image using Unicode braille characters.
func Write(source image.Image, options ...Option) error {
	// default options if none are passed in
	opts := &writeOpts{
		width:  100,
		writer: os.Stdout,
		invert: false,
	}

	// apply options over the defaults
	for _, opt := range options {
		opt(opts)
	}

	// scale the image to have exactly the number of pixels wide as there are
	// braille dots.
	img := imaging.Resize(source, opts.width*2, 0, imaging.Lanczos)

	max := img.Bounds().Max
	points := [8]bool{}

	calc := func(x, y int) bool {
		var res bool

		r, g, b, alpha := img.At(x, y).RGBA()
		if r > wc && g > wc && b > wc {
			res = false
		} else {
			res = alpha >= halpha
		}

		if opts.invert {
			return !res
		}

		return res
	}

	for y := 0; y < max.Y; y += 4 {
		for x := 0; x < max.X; x += 2 {
			// create a braile char for each 2x4 pixels
			points[0] = calc(x, y)
			points[1] = calc(x+1, y)
			points[2] = calc(x, y+1)
			points[3] = calc(x+1, y+1)
			points[4] = calc(x, y+2)
			points[5] = calc(x+1, y+2)
			points[6] = calc(x, y+3)
			points[7] = calc(x+1, y+3)

			s := braille.Get(points)
			fmt.Fprint(opts.writer, s)
		}
		fmt.Fprintln(opts.writer)
	}
	return nil
}
