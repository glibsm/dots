package braille

import (
	"testing"
)

const (
	X = true
	O = false
)

func TestBrailleConversion(t *testing.T) {
	tts := []struct {
		input  [size]bool
		expect string
	}{
		{
			input: [size]bool{
				X, O,
				O, X,
				X, O,
				O, X,
			},
			expect: "⢕",
		},
		{
			input: [size]bool{
				X, O,
				X, O,
				X, O,
				X, O,
			},
			expect: "⡇",
		},
		{
			input: [size]bool{
				O, O,
				X, X,
				X, X,
				O, O,
			},
			expect: "⠶",
		},
	}
	for _, tt := range tts {
		res := Get(tt.input)
		if res != tt.expect {
			t.Error("expected", tt.expect, "for", tt.input, "got", res)
		}
	}
}
