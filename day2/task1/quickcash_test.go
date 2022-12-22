package task1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCashFromSavingsAccount(t *testing.T) {

	sa := &SavingsAccount{
		balance: 5000,
	}
	cca := &CreditCardAccount{
		creditLimit: 5000,
	}

	qc := QuickCash{
		sa,
		cca,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, sa.GetIdentifier(), accType)
}

func TestGetCashFromCreditCardAccount(t *testing.T) {

	sa := &SavingsAccount{
		balance: 100,
	}
	cca := &CreditCardAccount{
		creditLimit: 5000,
	}

	qc := QuickCash{
		sa,
		cca,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, cca.GetIdentifier(), accType)
}
