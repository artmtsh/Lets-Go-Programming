package myCalc

import "fmt"

func Calculation(num1 float32, operation string, num2 float32) float32 {
	switch operation {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
	return 0.0
}
