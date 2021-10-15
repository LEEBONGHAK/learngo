package main

import (
	"fmt"

	"github.com/LEEBONGHAK/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("bonghak")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
