package main

import (
	"fmt"
)

func Task1(input []int) int {

	var Counter int

	for i := 0; i < len(input); i++ {
		Counter++
	}

	if input[len(input)-1] < Counter {
		return Counter - 1
	}

	if input[len(input)-1] == Counter {
		return Counter + 1
	}
	return Counter
}

func main() {
	fmt.Println(Task1([]int{1, 3, 6, 4, 1, 2}))
	fmt.Println(Task1([]int{1, 2, 3}))
	fmt.Println(Task1([]int{-1, -3}))
}
