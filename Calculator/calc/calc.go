package calc

import "fmt"

//Add two numbers
func Add(num1, num2 *float64) *float64 {
	result := (*num1) + (*num2)
	return &result
}

//Subtract two numbers
func Subtract(num1, num2 *float64) *float64 {
	result := (*num2) - (*num2)
	return &result
}

//Multiply two numbers
func Multiply(num1, num2 *float64) *float64 {
	result := (*num2) * (*num2)
	return &result
}

//Divide two numbers
func Divide(num1, num2 *float64) (*float64, error) {
	if *num2 == 0 {
		return nil, fmt.Errorf("Can't devide by zero")
	}
	result := (*num2) / (*num2)
	return &result, nil
}
