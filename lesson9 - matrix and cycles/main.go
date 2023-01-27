package main

import (
	"fmt"
)

func main() {
	matrix := [][]int {}
	matrix = make([][]int, 10)
	counter := 0
	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			counter++
			matrix[x][y] = counter;
		}
	}
	print(matrix)

	arr := []string {"1", "2", "3"}
	// спец синтаксис прохода по массиву
	for i := range arr {
		fmt.Println(i, arr[i])
	}

	count := 0

	for {
		if (count == 100) {
			break
		}
		count++
		fmt.Println(count)
	}
}

func print(arr [][]int)  {
	fmt.Println(arr)
}
