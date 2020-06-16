package dots

import (
	"bytes"
	"image"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestShapes(t *testing.T) {
	tcs := []struct {
		name     string
		file     string
		expected string
	}{
		{
			file: "cross.png",
			expected: `
⠑⢄⠀⠀⠀⠀⡠⠊
⠀⠀⠑⢄⡠⠊⠀⠀
⠀⠀⡠⠊⠑⢄⠀⠀
⡠⠊⠀⠀⠀⠀⠑⢄`,
		},
		{
			file: "square.png",
			expected: `
⡏⠉⠉⠉⠉⠉⠉⢹
⡇⠀⠀⠀⠀⠀⠀⢸
⡇⠀⠀⠀⠀⠀⠀⢸
⣇⣀⣀⣀⣀⣀⣀⣸`,
		},
		{
			file: "checker.png",
			expected: `
⠛⣤⠛⣤⠛⣤⠛⣤
⠛⣤⠛⣤⠛⣤⠛⣤
⠛⣤⠛⣤⠛⣤⠛⣤
⠛⣤⠛⣤⠛⣤⠛⣤`,
		},
	}

	var buf bytes.Buffer
	for _, tc := range tcs {
		buf.Reset()

		f, err := os.Open(filepath.Join("testdata", tc.file))
		if err != nil {
			t.Fatal("failed to read image file", err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			t.Fatal("failed to decode image", err)
		}

		if err := Write(img, Writer(&buf), Width(8)); err != nil {
			t.Fatal("failed to write image bytes", err)
		}

		if strings.TrimSpace(buf.String()) != strings.TrimSpace(tc.expected) {
			t.Fail()
			t.Logf("expected:%s\ngot:\n%s\n", tc.expected, buf.String())
		}
	}
}
