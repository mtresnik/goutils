package goutils

import (
	"image/color"
	"image/draw"
)

func DrawLine(img draw.Image, x0, y0, x1, y1 int, color color.Color) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx := 1
	if x0 > x1 {
		sx = -1
	}
	sy := 1
	if y0 > y1 {
		sy = -1
	}
	err := dx - dy

	for {
		img.Set(x0, y0, color)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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
