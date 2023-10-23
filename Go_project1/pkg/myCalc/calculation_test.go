package myCalc

import (
	"testing"
)

func TestAddition(t *testing.T) {
	var num1 float32 = 6
	var num2 float32 = 5
	var operation = "+"
	var result = Calculation(num1, operation, num2)
	var expected float32 = 11.0
	if result != expected {
		t.Errorf("Addition is incorrect")
	}
}

func TestSubtraction(t *testing.T) {
	var num1 float32 = 6
	var num2 float32 = 5
	var operation = "-"
	var result = Calculation(num1, operation, num2)
	var expected float32 = 1.0
	if result != expected {
		t.Errorf("Subtraction is incorrect")
	}
}
func TestMultiplication(t *testing.T) {
	var num1 float32 = 6
	var num2 float32 = 5
	var operation = "*"
	var result = Calculation(num1, operation, num2)
	var expected float32 = 30.0
	if result != expected {
		t.Errorf("Multiplication is incorrect")
	}
}
func TestDivision(t *testing.T) {
	var num1 float32 = 6
	var num2 float32 = 5
	var operation = "/"
	var result = Calculation(num1, operation, num2)
	var expected float32 = 1.2
	if result != expected {
		t.Errorf("Division is incorrect")
	}
}

func TestDivisionByZero(t *testing.T) {
	// Подготовка данных для деления на ноль
	numerator := 5.0
	denominator := 0.0

	// Защит блок для проверки деления на ноль
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Деление не вызвало ожидаемой паники")
		}
	}()

	// Попытка деления на ноль, которая должна вызвать панику
	result := numerator / denominator
	_ = result
}
