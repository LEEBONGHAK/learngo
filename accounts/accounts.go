package accounts

// account struct
type account struct {
	// 소문자일 경우 private, 첫글자가 대문자면 public
	owner   string
	balance int
}

// struct의 수정을 막기 위해서 struct를 만드는 함수
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// Deposit method
func (accountName account) Deposit(amount int) {
	accountName.balance += amount
}

// Balance of account
func (accountName account) Balance() int {
	return accountName.balance
}
