package accounts

import "errors"

// account struct
type account struct {
	// 소문자일 경우 private, 첫글자가 대문자면 public
	owner   string
	balance int
}

// error's name should be started 'err'
var errNoMoney = errors.New("Can't withdraw")

// struct의 수정을 막기 위해서 struct를 만드는 함수
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// Deposit method
func (accountName *account) Deposit(amount int) {
	accountName.balance += amount
}

// Withdraw method
func (accountName *account) Withdraw(amount int) error {
	if accountName.balance < amount {
		return errNoMoney
	}
	accountName.balance -= amount
	// nill : like none or null
	return nil
}

// Balance of account
func (accountName account) Balance() int {
	return accountName.balance
}
