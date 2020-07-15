package image

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

const (
	dx = 500
	dy = 200
)

func TestWritePng(t *testing.T) {
	// react
	react := image.Rect(0, 0, dx, dy)
	// alpha
	alpha := image.NewAlpha(react)

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			alpha.Set(x, y, color.Alpha{A: uint8(x % 256)})
		}
	}
	file, err := os.Create("test.png")
	if err != nil {
		t.Fatal(err)
	}
	png.Encode(file, alpha)
	os.Remove("test.png")
}

func TestReadPng(t *testing.T) {
	pngRead, err := os.Open("testdata/test.png")
	defer pngRead.Close()
	if err != nil {
		t.Fatal(err)
	}
	pngImg, err := png.Decode(pngRead)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pngImg.Bounds())
}
