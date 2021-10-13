package main

// module을 export하고 싶으면 함수 이름의 첫글자를 대문자로 해야함
import (
	"fmt"

	"github.com/LEEBONGHAK/learngo/something"
)

func main() {
	// fmt : formatting을 관리하는 package
	fmt.Println("Hello World")
	something.SayHello()

	// const: unchangeable untype variable
	const name = "bonghak"
	// variable name 뒤에 type 결정 가능, string, bool, ...
	const name2 string = "lee"

	// var: changeable untype variable, const와 마찬가지로 뒤에 type 결정 가능
	var a string = "nico"

	// func 안에만 가능, Go가 type을 찾아주며 한번 type을 결정해주면 해당 type으로만 사용해야함
	b := "lee" // var b string = "lee"
}
