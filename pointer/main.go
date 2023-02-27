package main

import "fmt"

type Rect struct {
	w, h float64
}

func (r Rect) Square() float64 {
	return r.w * r.h
}

// Can modify the value of which the reciever points to
func (r *Rect) Scale(f float64) {
	r.w = r.w * f
	r.h = r.h * f
}

// Can not change the values of the passed-in parameters
func (r Rect) Scale1(f float64) {
	r.w = r.w * f
	r.h = r.h * f
}

func ScaleFunc(r *Rect, f float64) {
	r.w = r.w * f
	r.h = r.h * f
}

func main() {
	var r = Rect{9, 10}
	var r1 = Rect{9, 10}
	r.Scale(10)
	r1.Scale1(10)

	fmt.Println(r.Square())
	fmt.Println(r1.Square())

	var r2, r3 = Rect{9, 10}, Rect{9, 10}
	r2.Scale(10)
	ScaleFunc(&r3, 10)

	// r2.Scale(10) OK
	// ScaleFunc(r2, 10) compiler error

}
