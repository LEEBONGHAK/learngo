package main

import "fmt"

func main() {
	// map : python이나 JS의 object 같은 것
	test := map[string]string{"name": "bonghak", "age": "12"}
	fmt.Println(test)
	for key, value := range test {
		fmt.Println(key, value)
	}
}
