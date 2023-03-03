package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Pix  []int
	Rect image.Rectangle
}

func (p Image) Bounds() image.Rectangle {
	return p.Rect
}

func (p Image) At(x, y int) color.Color {
	return color.RGBA{255, 255, 255, 0}
}

func (p Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
