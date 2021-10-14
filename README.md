# Learn Go  
  
Learning Go and making web scrapper  
Lecture: [쉽고 빠른 Go 시작하기](https://nomadcoders.co/go-for-beginners)  
  
## Main Package and Imports
  
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

func repeatMe (words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("1", "2", "3", "4", "5") // output: [1, 2, 3, 4, 5]
}
```  
  
"naked" return: return할 값을 variable처럼 사용 가능하게 하는 것  
```
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upperName := lenAndUpper("bonghak")
	fmt.Println(totalLength, upperName)
}

```  
  
defer: function이 끝난 후 수행하는 line  
```
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")	// 함수가 끝나고 "I'm done" 출력
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upperName := lenAndUpper("bonghak")
	fmt.Println(totalLength, upperName)
}
```  
  
## for, range, ...args  
  
range : array에 loop을 적용할 수 있도록 해줌, index와 넘겨준 것을 던져줌  
```
package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0

	fmt.Println(numbers)
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}

	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5)
	fmt.Println(result)
}
```  
  
for loop를 하는 다른 방법  
```
package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0

	fmt.Println(numbers)
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
		fmt.Println(i, numbers[i])
	}

	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5)
	fmt.Println(result)
}
```  
  
## If  
  
기본적인 If 사용법  
```
func canIDrink(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}
```  
  
variable expression : if문에서 variable을 생성하는 방법  
variable expression은 if문에서만 사용 가능하고, 이 variable은 if문에서만 사용될 것이란 것을 읽는 사람이 쉽게 알 수 있음  
```
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
```  
  
## Switch  
  
일반적인 switch 사용법  
```
func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
```  
  
variable expression을 적용한 switch  
```
func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
```  
  
조건을 주면서 사용하는 것도 가능하다.  
```
func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}
```  
  
## Pointers  
  
&(variable name) : 메모리 주소를 보는 법  
```
package main

import "fmt"

func main() {
	a := 2
	b := 5
	fmt.Println(&a, &b)
}
```  
  
`*` : 메모리 주소에 저장되어 있는 데이터를 보는 법  
```
package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Println(*b)
}
```  
  
## Arrays and Slices  
  
array를 만드는 방법 : `(array name) := [(array size)](array type)){(elements)}`  
example : `names := [5]string{"a", "b", "c", "d", "e"}`  
  
slice : 기본적으로 array이지만 size 지정 없이 사용하는 것 (size가 자동으로 변화)  
example : `	numbers := []int{1, 2, 3, 4, 5}`  
  
append : array에 element를 추가하는 함수 `append((array name), (want to add element))`  
append는 array를 수정하는 것이 아니라 업데이트하여 return하는 것이기 때문에 아래와 같이 사용해야 한다.  
`numbers = append(numbers, 6)`  
  
## Maps  
  
