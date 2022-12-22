package main

import (
	"fmt"
)

func main() {
	var conunter = increment()
	fmt.Println(conunter())
	fmt.Println(conunter())
}

func increment() func() int {
	var count = 0
	return func() int {
		count++
		return count
	}
}
