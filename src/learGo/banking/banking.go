package banking

import (
	"errors"
	"fmt"
)

//Account Stuct
type Account struct {
	owner string
	balance int
}

var errNoMoney = errors.New("can't withdraw")

//NewAccount creates Account
func NewAccount(owner string) *Account{
	account:= Account{owner:owner, balance:0}
	return &account
}

//Deposit 리시버 a 위에서 주소를 반환했잖아 그러니까 값에 접근 하기위해서는 해당 값이라는걸 알려줘야 하는 것이 아닐까하는데 말이지?
func (a *Account) Deposit(amount int) {
	fmt.Println("amount : ",amount)
	a.balance += amount
}

//Balance receiver
func (a *Account) Balance() int{
	return a.balance
}

//Withdraw is
func (a *Account) Withdraw(amount int) error{
	if a.balance < amount{
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//ChangeOwner is..
func (a *Account) ChangeOwner(newOwner string){
	a.owner  = newOwner
}
//Owner is
func (a Account) Owner() string{
	//카피 본이 된다.
	return a.owner
}