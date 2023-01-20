package main

import (
	"fmt"
)

func main() {
	var arr = []int {1, 2} // всратый синтаксис для массивов
	fmt.Println(arr)
	fmt.Println(minNum(1, 2, 5, 3, -5, 2))
}

func minNum(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	var min = numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
	}

	return min
}
