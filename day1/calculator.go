package main

import (
	"errors"
	"math"
)

var errDivisionByZero = errors.New("Cannot divide by zero")

func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a float64, b float64) float64 {
	return a - b
}

func multiply(a float64, b float64) float64 {
	return a * b
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errDivisionByZero
	}
	return a / b, nil
}

func sin(a float64) float64 {
	return math.Sin(a)
}

func cos(a float64) float64 {
	return math.Cos(a)
}

func tan(a float64) float64 {
	return math.Tan(a)
}

func root(a float64) float64 {
	return math.Sqrt(a)
}
