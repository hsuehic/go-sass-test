package main

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i = i + 1 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Sum(s []int, c chan int) {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	c <- sum
}

func main() {
	go Say("Hello")
	Say("World")

	var s = []int{1, 2, 3, 4, 8, 8, 20, -8, 9}
	var c = make(chan int, 5)
	go Sum(s[:len(s)/2], c)
	go Sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("x: %v, y: %v, x+y: %v\n", x, y, x+y)
}
