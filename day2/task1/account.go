package task1

type SavingsAccount struct {
	balance float64
}

func (sa *SavingsAccount) WithDraw(amount float64) {
}

func (sa *SavingsAccount) CanWithDraw(amount float64) bool {
	return amount <= sa.balance
}

func (sa *SavingsAccount) GetIdentifier() string {
	return "SAVINGS"
}

type CreditCardAccount struct {
	creditLimit float64
}

func (cca *CreditCardAccount) WithDraw(amount float64) {
}

func (cca *CreditCardAccount) CanWithDraw(amount float64) bool {
	return amount <= cca.creditLimit
}

func (cca *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD"
}
