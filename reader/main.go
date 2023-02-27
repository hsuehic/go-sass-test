// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import (
	"fmt"
	"strings"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	a := strings.Repeat("A", len(b))
	copy(b, a)
	return len(b), nil
}

func main() {
	buffer := make([]byte, 20)
	reader := MyReader{}
	l, _ := reader.Read(buffer)
	if l > 0 {
		fmt.Println(string(buffer))
	}
}
