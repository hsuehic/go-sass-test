package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		fmt.Println(t.Value)
		Walk(t.Right, ch)
	}
}

func WalkTree(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go WalkTree(t1, ch1)
	go WalkTree(t2, ch2)
	// for x := range ch1 {
	// 	fmt.Println(x)
	// }
	// for y := range ch2 {
	// 	fmt.Println(y)
	// }

	for i := 0; i < 11; i++ {
		x, ok1 := <-ch1
		y, ok2 := <-ch2

		fmt.Printf("x: %v, y: %v, ok1: %v, ok2: %v\n", x, y, ok1, ok2)
		if x != y || ((!ok1 || !ok2) && ok1 != ok2) {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	t3 := tree.New(1)
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t3))
}
