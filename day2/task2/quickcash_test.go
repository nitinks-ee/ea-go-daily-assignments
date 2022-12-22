package task2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCashFromPaytmWallet(t *testing.T) {

	pw := &PaytmWallet{
		balance: 5000,
	}
	sa := &SavingsAccount{
		balance: 5000,
	}
	cca := &CreditCardAccount{
		creditLimit: 5000,
	}

	qc := QuickCash{
		[]Withdrawable{cca, sa, pw},
	}

	amt, accType, _ := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, pw.GetIdentifier(), accType)
}

func TestGetCashFromSavingsAccount(t *testing.T) {

	pw := &PaytmWallet{
		balance: 100,
	}
	sa := &SavingsAccount{
		balance: 5000,
	}
	cca := &CreditCardAccount{
		creditLimit: 5000,
	}

	qc := QuickCash{
		[]Withdrawable{cca, sa, pw},
	}

	amt, accType, _ := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, sa.GetIdentifier(), accType)
}

func TestGetCashFromCreditCardAccount(t *testing.T) {

	pw := &PaytmWallet{
		balance: 100,
	}
	sa := &SavingsAccount{
		balance: 100,
	}
	cca := &CreditCardAccount{
		creditLimit: 5000,
	}

	qc := QuickCash{
		[]Withdrawable{cca, sa, pw},
	}

	amt, accType, _ := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, cca.GetIdentifier(), accType)
}

func TestLimitExaustedError(t *testing.T) {

	pw := &PaytmWallet{
		balance: 100,
	}
	sa := &SavingsAccount{
		balance: 100,
	}
	cca := &CreditCardAccount{
		creditLimit: 100,
	}

	qc := QuickCash{
		[]Withdrawable{cca, sa, pw},
	}

	_, _, err := qc.getCash(500)
	assert.Equal(t, err, errLimitExceeded)
}
