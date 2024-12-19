package goutils

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

func DrawLine(img draw.Image, x0, y0, x1, y1 int, color color.Color, thickness int) {
	minX := min(x0, x1)
	minY := min(y0, y1)
	maxX := max(x0, x1)
	maxY := max(y0, y1)

	dx := x1 - x0
	dy := y1 - y0

	minX = minX - thickness/2
	maxX = maxX + thickness/2
	minY = minY - thickness/2
	maxY = maxY + thickness/2

	numPoints := int(math.Sqrt(float64(dx*dx + dy*dy)))

	for index := 0; index < numPoints; index++ {
		t := float64(index) / float64(numPoints)
		FillCircle(img, int(t*float64(dx))+x0, int(t*float64(dy))+y0, thickness/2, color)
	}
}

func DrawRectangle(img draw.Image, x0, y0, width, height int, color color.Color) {
	for x := x0; x < x0+width; x++ {
		img.Set(x, y0, color)
		img.Set(x, y0+height-1, color)
	}
	for y := y0; y < y0+height; y++ {
		img.Set(x0, y, color)
		img.Set(x0+width-1, y, color)
	}
}

func FillCircle(img draw.Image, x0, y0, radius int, color color.Color) {
	minX := x0 - radius
	maxX := x0 + radius
	minY := y0 - radius
	maxY := y0 + radius
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if x < 0 || y < 0 || x >= img.Bounds().Max.X || y >= img.Bounds().Max.Y {
				continue
			}
			if (x-x0)*(x-x0)+(y-y0)*(y-y0) <= radius*radius {
				img.Set(x, y, color)
			}
		}
	}
}

func FillRectangle(img draw.Image, x0, y0, width, height int, color color.Color) {
	for x := x0; x < x0+width; x++ {
		for y := y0; y < y0+height; y++ {
			if x < 0 || y < 0 || x >= img.Bounds().Max.X || y >= img.Bounds().Max.Y {
				continue
			}
			img.Set(x, y, color)
		}
	}
}

func HashColor(c color.Color) int64 {
	r, g, b, a := c.RGBA()
	values := []float64{float64(r), float64(g), float64(b), float64(a)}
	return HashFloats(values...)
}

func GetColors(img *image.RGBA) []color.Color {
	uniqueColors := make(map[int64]bool)
	retArray := make([]color.Color, 0)
	for x := 0; x < img.Rect.Dx(); x++ {
		for y := 0; y < img.Rect.Dy(); y++ {
			col := img.At(x, y)
			hash := HashColor(col)
			if _, ok := uniqueColors[hash]; !ok {
				uniqueColors[hash] = true
				retArray = append(retArray, col)
			}
		}
	}
	return retArray
}

func ConvertImageToPaletted(img *image.RGBA, colors ...color.RGBA) *image.Paletted {
	palette := make([]color.Color, len(colors))
	if len(palette) > 0 {
		for i, c := range colors {
			palette[i] = c
		}
	} else {
		palette = GetColors(img)
	}
	retImage := image.NewPaletted(img.Rect, palette)
	for x := 0; x < img.Rect.Dx(); x++ {
		for y := 0; y < img.Rect.Dy(); y++ {
			retImage.Set(x, y, img.At(x, y))
		}
	}
	return retImage
}
