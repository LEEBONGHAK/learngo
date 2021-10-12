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
}
