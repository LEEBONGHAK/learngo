package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0
	fmt.Println(numbers)
	// range: array에 loop을 적용할 수 있도록 해줌, index와 넘겨준 것을 던져줌
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	fmt.Println()

	// for loop를 하는 다른 방법
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5)
	fmt.Println(result)
}
