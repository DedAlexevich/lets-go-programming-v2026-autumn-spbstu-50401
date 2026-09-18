package main

import "fmt"

func main() {
	var a, b int
	var op rune
	_, efirst := fmt.Scan(&a)
	_, esecond := fmt.Scan(&b)
	_, err := fmt.Scanf("%c\n", &op)

	if efirst != nil {
		fmt.Println("Invalid first operand")
		return
	} else if esecond != nil {
		fmt.Println("Invalid second operand")
		return
	} else if err != nil {
		fmt.Println("Invalid operation")
	}

	switch op {
	case '+':
		fmt.Println(a + b)
	case '-':
		fmt.Println(a - b)
	case '*':
		fmt.Println(a * b)
	case '/':
		if b == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Invalid operation")
		return
	}
}
