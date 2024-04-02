package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the first number: ")
	num1Str, _ := reader.ReadString('\n')
	num1, err := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)
	if err != nil {
		fmt.Println("Error parsing the first number:", err)
		return
	}

	fmt.Print("Enter the operator (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	fmt.Print("Enter the second number: ")
	num2Str, _ := reader.ReadString('\n')
	num2, err := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)
	if err != nil {
		fmt.Println("Error parsing the second number:", err)
		return
	}

	result, err := calculate(num1, operator, num2)
	if err != nil {
		fmt.Println("Error calculating result:", err)
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}

func calculate(num1 float64, operator string, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
}
