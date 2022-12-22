package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	acc.Withdraw(200)
	assert.Equal(t, float64(300), acc.GetBalance())
}

func TestErrorWithdrawal(t *testing.T) {
	acc := Account{balance: 100}
	err := acc.Withdraw(200)
	sfError := err.(ShortageFundsError)
	var balance float64 = 100
	assert.Equal(t, sfError.balance, balance)
	assert.Equal(t, sfError.short, balance)
	assert.Equal(t, sfError.reason, "Insufficient balance")
}
