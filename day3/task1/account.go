package account

import "fmt"

type ShortageFundsError struct {
	balance float64
	short   float64
	reason  string
}

func (sf ShortageFundsError) Error() string {
	return fmt.Sprintf("%s, Current Balance- %f, Shortage funds- %f", sf.reason, sf.balance, sf.short)
}

type Account struct {
	balance float64
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64) error {
	if amount > acc.balance {
		return ShortageFundsError{acc.balance, amount - acc.balance, "Insufficient balance"}
	}
	acc.balance -= amount
	return nil
}
