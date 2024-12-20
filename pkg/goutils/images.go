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

func FloatsToColors(floats [][]float64) []color.Color {
	ret := make([]color.Color, len(floats))
	for i, f := range floats {
		ret[i] = FloatToColor(f)
	}
	return ret
}

func FloatToColor(point []float64) color.Color {
	r := uint8(point[0])
	g := uint8(point[1])
	b := uint8(point[2])
	a := uint8(point[3])
	return color.RGBA{R: r, G: g, B: b, A: a}
}

func ColorsToFloats(colors []color.Color) [][]float64 {
	ret := make([][]float64, len(colors))
	for i, c := range colors {
		ret[i] = ColorToFloats(c)
	}
	return ret
}

func ColorToFloats(c color.Color) []float64 {
	r, g, b, a := c.RGBA()
	return []float64{float64(r), float64(g), float64(b), float64(a)}
}

func HashColor(c color.Color) int64 {
	return HashFloats(ColorToFloats(c)...)
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
		imgColors := GetColors(img)
		if len(imgColors) > 255 {
			palette = FloatsToColors(kMeansFloats(255, ColorsToFloats(imgColors)...))
		} else {
			palette = imgColors
		}
	}
	retImage := image.NewPaletted(img.Rect, palette)
	for x := 0; x < img.Rect.Dx(); x++ {
		for y := 0; y < img.Rect.Dy(); y++ {
			retImage.Set(x, y, img.At(x, y))
		}
	}
	return retImage
}

var COLOR_RED = color.RGBA{255, 0, 0, 255}
var COLOR_YELLOW = color.RGBA{255, 255, 0, 255}
var COLOR_BLUE = color.RGBA{0, 0, 255, 255}
var COLOR_MAGENTA = color.RGBA{255, 0, 255, 255}
var COLOR_CYAN = color.RGBA{0, 255, 255, 255}
var COLOR_WHITE = color.RGBA{255, 255, 255, 255}
var COLOR_BLACK = color.RGBA{0, 0, 0, 255}
var COLOR_GRAY = color.RGBA{128, 128, 128, 255}
var COLOR_LIGHT_GRAY = color.RGBA{192, 192, 192, 255}
var COLOR_DARK_GRAY = color.RGBA{64, 64, 64, 255}
var COLOR_GREEN = color.RGBA{0, 255, 0, 255}
var COLOR_ORANGE = color.RGBA{255, 165, 0, 255}
var COLOR_PINK = color.RGBA{255, 192, 203, 255}
var COLOR_PURPLE = color.RGBA{128, 0, 128, 255}
var COLOR_BROWN = color.RGBA{165, 42, 42, 255}
var COLOR_LIGHT_GREEN = color.RGBA{144, 238, 144, 255}
var COLOR_LIGHT_BLUE = color.RGBA{173, 216, 230, 255}
var COLOR_LIGHT_YELLOW = color.RGBA{255, 255, 224, 255}
var COLOR_LIGHT_ORANGE = color.RGBA{255, 204, 102, 255}

func GradientGreenToRed(scalar float64) color.RGBA {
	return Gradient(scalar, COLOR_GREEN, COLOR_YELLOW, COLOR_RED)
}

func Gradient(pScalar float64, firstColor color.RGBA, colors ...color.RGBA) color.RGBA {
	if len(colors) == 0 {
		return firstColor
	}
	scalar := pScalar
	if scalar < 0 {
		scalar = 0
	}
	if scalar > 1 {
		scalar = 1
	}
	allColors := append([]color.RGBA{firstColor}, colors...)
	var colorToScalar []Tuple
	maxScalar := float64(0)
	for i, c := range allColors {
		colorToScalar = append(colorToScalar, Pair(HashColor(c), float64(i)/float64(len(allColors)-1)))
		maxScalar = max(maxScalar, float64(i)/float64(len(allColors))-1)
	}
	minIndex := 0
	maxIndex := len(colorToScalar) - 1
	minRange := colorToScalar[minIndex]
	maxRange := colorToScalar[maxIndex]
	for i, tuple := range colorToScalar {
		_, keyok := tuple[0].(int64)
		value, valueok := tuple[1].(float64)
		if keyok && valueok {
			if scalar > value {
				minIndex = i
				minRange = colorToScalar[minIndex]
			} else if scalar < value {
				maxIndex = i
				maxRange = colorToScalar[maxIndex]
				break
			} else {
				return allColors[i]
			}
		}
	}
	minScalar, minOk := minRange[1].(float64)
	maxScalar, maxOk := maxRange[1].(float64)
	if minOk && maxOk {
		t := (scalar - minScalar) / (maxScalar - minScalar)
		minColor := allColors[minIndex]
		maxColor := allColors[maxIndex]
		retColor := color.RGBA{
			R: uint8(float64(minColor.R)*(1.0-t) + float64(maxColor.R)*t),
			G: uint8(float64(minColor.G)*(1.0-t) + float64(maxColor.G)*t),
			B: uint8(float64(minColor.B)*(1.0-t) + float64(maxColor.B)*t),
			A: uint8(float64(minColor.A)*(1.0-t) + float64(maxColor.A)*t),
		}
		return retColor
	}
	return allColors[minIndex]
}

func kMeansFloats(n int, points ...[]float64) [][]float64 {
	if n <= 0 || len(points) == 0 {
		return nil
	}

	centroids := make([][]float64, n)
	for i := range centroids {
		centroids[i] = points[i%len(points)]
	}

	distance := func(a, b []float64) float64 {
		sum := 0.0
		for i := range a {
			diff := a[i] - b[i]
			sum += diff * diff
		}
		return sum
	}

	assignments := make([]int, len(points))
	for {
		changed := false

		for i, point := range points {
			closest := 0
			minDist := distance(point, centroids[0])
			for j := 1; j < n; j++ {
				dist := distance(point, centroids[j])
				if dist < minDist {
					closest = j
					minDist = dist
				}
			}
			if assignments[i] != closest {
				assignments[i] = closest
				changed = true
			}
		}

		if !changed {
			break
		}

		counts := make([]int, n)
		newCentroids := make([][]float64, n)
		for i := range newCentroids {
			newCentroids[i] = make([]float64, len(points[0]))
		}
		for i, point := range points {
			cluster := assignments[i]
			for j := range point {
				newCentroids[cluster][j] += point[j]
			}
			counts[cluster]++
		}
		for i := range newCentroids {
			if counts[i] == 0 {
				continue
			}
			for j := range newCentroids[i] {
				newCentroids[i][j] /= float64(counts[i])
			}
		}
		centroids = newCentroids
	}

	return centroids
}
