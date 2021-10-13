# Learn Go  
  
Learning Go and making web scrapper  
Lecture: [쉽고 빠른 Go 시작하기](https://nomadcoders.co/go-for-beginners)  
  
## Main Package and 
  
package main: Go의 기본 package  
fmt: formatting을 관리하는 module (module을 사용하려면 import)
  
module을 export하고 싶으면 함수 이름의 첫글자를 대문자로 해야함  
```
package something

import "fmt"

func sayBye() {   // It doesn't work as module
	fmt.Println("Bye")
}

func SayHello() {  // It work as module
	fmt.Println("Hello")
}
```  
  
## Variables and Constants  
  
const: unchangeable untype variable  
`const name = "bonghak"`  
variable name 뒤에 type 결정 가능, string, bool, int8, ...  
`const name string = "bonghak"`  
  
var: changeable untype variable, const와 마찬가지로 뒤에 type 결정 가능  
var을 사용한 variable들은 무조건 한번은 사용해야함 안하면 에러  
`var name string = "bonghak"`  
아래와 같이 축약 가능하지만 func 안에만 가능, Go가 type을 찾아주며 한번 type을 결정해주면 해당 type으로만 사용해야함  
`name := "bonghak" // var name string = "bonghak"`  
  
## Functions  
  
Go의 function은 `func`으로 시작  
argument와 return의 type을 지정해 줘야함  
```
func multiply(a int, b int) int { 
	return a * b
}
```  
argument들의 type이 모두 같다면 아래와 같이 축약 가능  
```
func multiply(a, b int) int {
	return a * b
}
```  
  
여러 개 return 가능하며, type의 종류도 여러가지 동시에 return 가능  
```
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
```  
return 값은 한 만큼 받아야 함  
```
totalLength, upperName := lenAndUpper("bonghak")
fmt.Println(totalLength, upperName)
```  
모든 return을 받지 않는 법: '_' 사용  
```
totalLength, _ := lenAndUpper("bonghak")
fmt.Println(totalLength)
```  
  
무제한의 argument들을 받는 방법  
```
package main

import (
	"fmt"
)

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("1", "2", "3", "4", "5") // output: [1, 2, 3, 4, 5]
}
```  