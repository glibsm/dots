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

		img := loadImg(t, tc.file)
		if err := Write(img, Writer(&buf), Width(8)); err != nil {
			t.Fatal("failed to write image bytes", err)
		}

		if strings.TrimSpace(buf.String()) != strings.TrimSpace(tc.expected) {
			t.Fail()
			t.Logf("expected:%s\ngot:\n%s\n", tc.expected, buf.String())
		}
	}
}

func TestOptions(t *testing.T) {
	t.Run("invert", func(t *testing.T) {
		var buf bytes.Buffer
		img := loadImg(t, "cross.png")
		if err := Write(img, Writer(&buf), Width(8), Invert()); err != nil {
			t.Fatal("failed to write image bytes", err)
		}
		expected := `
⣮⡻⣿⣿⣿⣿⢟⣵
⣿⣿⣮⡻⢟⣵⣿⣿
⣿⣿⢟⣵⣮⡻⣿⣿
⢟⣵⣿⣿⣿⣿⣮⡻`
		got := buf.String()
		if strings.TrimSpace(expected) != strings.TrimSpace(got) {
			t.Fatalf("expected\n%s\ngot\n%s\n", expected, got)
		}
	})
}

func loadImg(t *testing.T, path string) image.Image {
	f, err := os.Open(filepath.Join("testdata", path))
	if err != nil {
		t.Fatal("failed to read image file", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		t.Fatal("failed to decode image", err)
	}

	return img
}
