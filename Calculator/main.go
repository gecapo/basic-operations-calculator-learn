package main

import (
	"fmt"
	"os"
	"strconv"

	"./calc"
)

func main() {
getOperation:
	var num1 string
	fmt.Print("Please input the first number: ")
	fmt.Scanln(&num1)
	firstNumber := stringToFloat64(num1)

	var operation string
	fmt.Print("Please select an operation: +, -, *, / : ")
	fmt.Scanln(&operation)

	var num2 string
	fmt.Print("Please input the second number: ")
	fmt.Scanln(&num2)
	secondNumber := stringToFloat64(num2)

	switch operation {
	case "+":
		fmt.Print("Result: ")
		fmt.Println(*calc.Add(firstNumber, secondNumber))
	case "-":
		fmt.Print("Result: ")
		fmt.Println(*calc.Subtract(firstNumber, secondNumber))

	case "*":
		fmt.Print("Result: ")
		fmt.Println(*calc.Multiply(firstNumber, secondNumber))

	case "/":
		fmt.Print("Result: ")
		fmt.Println(*calc.Divide(firstNumber, secondNumber))

	default:
		fmt.Println("Invalid operation selected. Please try again!")
		goto getOperation
	}
}

func stringToFloat64(str string) *float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return &f
}
