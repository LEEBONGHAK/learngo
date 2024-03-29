package main

import (
	"fmt"

	"github.com/LEEBONGHAK/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("bonghak")
	account.Deposit(10)
	fmt.Println(account.Balance(), account.Owner())
	err := account.Withdraw(20)
	if err != nil {
		// Fatall()n : 프로그램 종료 함수
		//log.Fatalln(err)
		fmt.Println(err)
	}
	account.ChangeOwner("Lee")
	fmt.Println(account.Balance(), account.Owner())

	fmt.Println(account)
}
