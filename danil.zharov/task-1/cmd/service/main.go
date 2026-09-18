package main

import "fmt"

func main() {
	var a, b int
	var operation string
	n, err := fmt.Scan(&a, &b, &operation)
	if err != nil && n == 0 {
		fmt.Println("Invalid first operand")
		return
	} else if err != nil && n == 1 {
		fmt.Println("Invalid second operand")
		return
	}

	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
