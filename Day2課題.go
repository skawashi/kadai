package main

import (
	"crypto/x509"
	"fmt"
)

func subtraction(x float64, y float64) float64 {
	ans := x - y
	return ans
}

func multiplication(x float64, y float64) float64 {
	ans := x * y
	return ans
}

func division(x float64, y float64) float64 {
	ans := x / y
	return ans
}

func calculation(operater string, x float64, y float64) float64 {
	var ans float64 = 0

	switch operater {
	case "+":
		ans = x + y

	case "-":
		ans = x - y

	case "*":
		ans = x * y

	case "/":
		ans = x / y
	}
	return ans
}

type Turtle struct {
	name string
	x	float64
	y	float64
	a	float64
}

func (t)

func main() {
	var num1 float64 = 12.5
	var num2 float64 = 2.5

	fmt.Println(subtraction(num1, num2))
	fmt.Println(multiplication(num1, num2))
	fmt.Println(division(num1, num2))
	fmt.Println(calculation("+", num1, num2))
}
