package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	if add(12.45, 20.22) != 32.67 {
		t.Fail()
	}
}

func TestSubtractTwoNumbers(t *testing.T) {
	if subtract(20.22, 12.45) != 7.77 {
		t.Fail()
	}
}

func TestMultiplyTwoNumbers(t *testing.T) {
	if multiply(12, 20) != 240 {
		t.Fail()
	}
}

func TestDivideTwoNumbers(t *testing.T) {
	result, _ := divide(10, 4)
	if result != 2.5 {
		t.Fail()
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(10, 0)
	if err != errDivisionByZero {
		t.Fail()
	}
}

func TestSineOfNumber(t *testing.T) {
	if sin(0) != 0 && sin(90) != 1 {
		t.Fail()
	}
}

func TestCosineOfNumber(t *testing.T) {
	if cos(0) != 1 && cos(90) != 0 {
		t.Fail()
	}
}

func TestTangentOfNumber(t *testing.T) {
	if tan(0) != 0 && tan(45) != 1 {
		t.Fail()
	}
}

func TestRootNumber(t *testing.T) {
	if root(4) != 2 && root(100) != 10 {
		t.Fail()
	}
}
