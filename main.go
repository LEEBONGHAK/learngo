package main

import (
	"fmt"

	"github.com/LEEBONGHAK/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("bonghak")
	fmt.Println(account)
}
