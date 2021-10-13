package main

// module을 export하고 싶으면 함수 이름의 첫글자를 대문자로 해야함
import (
	"fmt"
	"strings"

	"github.com/LEEBONGHAK/learngo/something"
)

func multiply(a int, b int) int { // func multiply(a, b int) int {
	return a * b
}

// Go에서는 여러개의 return 가능
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 무제한의 argument들을 받는 방법
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// fmt : formatting을 관리하는 package
	fmt.Println("Hello World")
	something.SayHello()

	// const: unchangeable untype variable
	const name = "bonghak"
	// variable name 뒤에 type 결정 가능, string, bool, int8, ...
	const name2 string = "lee"

	// var: changeable untype variable, const와 마찬가지로 뒤에 type 결정 가능
	// var을 사용한 variable들은 무조건 한번은 사용해야함 안하면 에러
	var a string = "nico"
	fmt.Println(a)

	// func 안에만 가능, Go가 type을 찾아주며 한번 type을 결정해주면 해당 type으로만 사용해야함
	b := "lee" // var b string = "lee"
	fmt.Println(b)

	fmt.Println(multiply(2, 2))

	// return하는 개수만큼 받아야함
	totalLength, upperName := lenAndUpper("bonghak")
	fmt.Println(totalLength, upperName)

	// 모든 return을 받지않는 법
	totalLength2, _ := lenAndUpper("bonghak")
	fmt.Println(totalLength2)

	repeatMe("1", "2", "3", "4", "5") // [1, 2, 3, 4, 5]
}
