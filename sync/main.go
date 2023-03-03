package main

import (
	"fmt"
	"sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	fmt.Println("Inside my goroutine")
}

func main() {
	fmt.Println("Hello World")

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	waitgroup.Wait()

	fmt.Println("Finished Execution")
}
