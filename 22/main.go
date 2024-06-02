package main

import (
	"fmt"
	"math/big"
)

type Calculator struct {
	a *big.Int
	b *big.Int
}

func NewCalculator(a, b *big.Int) *Calculator {
	return &Calculator{
		a: a,
		b: b,
	}
}

// Add stored values
func (calc *Calculator) Add() *big.Int {
	return big.NewInt(0).Add(calc.a, calc.b)
}

// Subtract stored values
func (calc *Calculator) Subtract() *big.Int {
	return big.NewInt(0).Sub(calc.a, calc.b)
}

// Multiply stored values
func (calc *Calculator) Multiply() *big.Int {
	return big.NewInt(0).Mul(calc.a, calc.b)
}

// Divide stored values
func (calc *Calculator) Divide() *big.Int {
	return big.NewInt(0).Div(calc.a, calc.b)
}

func main() {
	// Init big values
	a := big.NewInt(102983761298)
	b := big.NewInt(102938190)

	// Create calculator
	calc := NewCalculator(a, b)

	// Perform operations
	add := calc.Add()
	fmt.Println("add:", add)

	sub := calc.Subtract()
	fmt.Println("subtract:", sub)

	mul := calc.Multiply()
	fmt.Println("multiply:", mul)

	div := calc.Divide()
	fmt.Println("divide:", div)

}
