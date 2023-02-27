package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var r map[string]int = make(map[string]int)
	arr := strings.Fields(s)
	for _, str := range arr {
		v, ok := r[str]
		if ok {
			r[str] = v + 1
		} else {
			r[str] = 1
		}
	}
	return r
}

func main() {
	wc.Test(WordCount)
}
