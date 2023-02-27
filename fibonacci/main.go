package main

import "fmt"

func fibonacci() func() int {
	x := 0
	y := 1
	c := 0
	return func() int {
		var v int
		if c == 0 {
			v = x
		} else if c == 1 {
			v = y
		} else {
			v = x + y
			x = y
			y = v
		}
		c++
		return v
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; {
		fmt.Println(f())
		i += 1
	}
}
