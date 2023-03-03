package main

import (
	"fmt"
)

func Read(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Printf("Read: %v\n", i)
	}
	close(c)
}

func main() {

	fmt.Println("Channel without buffer")
	c1 := make(chan int)
	go Read(c1)
	for i := range c1 {
		fmt.Printf("Comsume: %v\n", i)
	}

	fmt.Println("Channel with buffer")
	c := make(chan int, 10)
	go Read(c)
	for i := range c {
		fmt.Printf("Comsume: %v\n", i)
	}

}
