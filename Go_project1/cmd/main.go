package main

import (
	"GO_PROJECT1/pkg/myCalc"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Введите первое число: ")
	var num1Str string
	fmt.Fscan(os.Stdin, &num1Str)
	num1, err1 := strconv.ParseFloat(num1Str, 32)
	if err1 != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}
	fmt.Print("Выберите операцию (+, -, *, /): ")
	var operation string
	fmt.Fscan(os.Stdin, &operation)
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		return
	}
	fmt.Print("Введите второе число: ")
	var num2Str string
	fmt.Fscan(os.Stdin, &num2Str)
	num2, err2 := strconv.ParseFloat(num2Str, 32)
	if err2 != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}
	fmt.Printf("Результат %.2f %s %.2f = %.2f", num1, operation, num2, myCalc.Calculation(float32(num1), operation, float32(num2)))
}
