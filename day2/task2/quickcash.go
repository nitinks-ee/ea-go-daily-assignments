package task2

import "errors"

var errLimitExceeded = errors.New("All accounts limit exceeded")

type QuickCash struct {
	accounts []Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string, error) {
	for i := len(qc.accounts) - 1; i >= 0; i-- {
		if qc.accounts[i].CanWithDraw(amount) {
			qc.accounts[i].WithDraw(amount)
			return amount, qc.accounts[i].GetIdentifier(), nil
		}
	}
	return 0, "", errLimitExceeded
}
