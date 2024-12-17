package goutils

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestDrawLine(t *testing.T) {
	imgRect := image.Rect(0, 0, 500, 500)
	img := image.NewRGBA(imgRect)

	x0 := 100
	y0 := 100
	x1 := 400
	y1 := 400
	lineColor := color.RGBA{255, 0, 0, 255} // Red color

	DrawLine(img, x0, y0, x1, y1, lineColor, 10)

	file, err := os.Create("TestDrawLine.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)

}

func TestFillCircle(t *testing.T) {
	imgRect := image.Rect(0, 0, 500, 500)
	img := image.NewRGBA(imgRect)

	x0 := 50
	y0 := 50
	lineColor := color.RGBA{255, 0, 0, 255} // Red color

	FillCircle(img, x0, y0, 100, lineColor)

	file, err := os.Create("TestFillCircle.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)

}

func TestFillRectangle(t *testing.T) {
	imgRect := image.Rect(0, 0, 500, 500)
	img := image.NewRGBA(imgRect)

	x0 := 50
	y0 := 50
	x1 := 150
	y1 := 150
	lineColor := color.RGBA{255, 0, 0, 255} // Red color

	FillRectangle(img, x0, y0, x1, y1, lineColor)

	file, err := os.Create("TestFillRectangle.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}

func TestConnectTheDots(t *testing.T) {
	imgRect := image.Rect(0, 0, 500, 500)
	img := image.NewRGBA(imgRect)

	x0 := 100
	y0 := 100
	x1 := 400
	y1 := 400
	pointColor := color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	}
	lineColor := color.RGBA{255, 0, 0, 255} // Red color

	FillCircle(img, x0, y0, 20, pointColor)
	FillCircle(img, x1, y1, 20, pointColor)
	DrawLine(img, x0, y0, x1, y1, lineColor, 10)

	file, err := os.Create("TestConnectTheDots.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}
