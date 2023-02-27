package main

import (
	"fmt"
	"math"
)

type Rect struct {
	w, h int
}

func (r Rect) area() int {
	return r.w * r.h
}

func area(r Rect) int {
	return r.w * r.h
}

type MyFloat float64

func (f MyFloat) abs() float64 {
	var v float64 = float64(f)
	if f < 0 {
		v = 0 - v
	}
	return v
}

func main() {
	var r Rect = Rect{10, 5}
	fmt.Println(r.area())

	var r1 Rect = Rect{18, 9}
	fmt.Println(area(r1))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.abs())
}
