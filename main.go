package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// struct : object와 비슷하면서 map보다 더 유연함 / structure 같은 것
	favFood := []string{"kimchi", "ramen"}
	person1 := person{name: "bonghak", age: 26, favFood: favFood}
	fmt.Println(person1)
}
