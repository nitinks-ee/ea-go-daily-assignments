package task2

type Withdrawable interface {
	CanWithDraw(amount float64) bool
	WithDraw(amount float64)
	GetIdentifier() string
}
