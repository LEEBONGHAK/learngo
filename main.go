package main

import "fmt"

func main() {
	names := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(names)
	// slice : 기본적으로 array이지만 size 없이 사용되는 것 (size가 자동으로 변화)
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6)
	fmt.Println(numbers)
}
