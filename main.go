package main

import (
	"fmt"
	"strings"
)

// "naked" return: return할 값을 variable로서 사용가능
// defer: function이 끝난 후 수행하는 line
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upperName := lenAndUpper("bonghak")
	fmt.Println(totalLength, upperName)
}
