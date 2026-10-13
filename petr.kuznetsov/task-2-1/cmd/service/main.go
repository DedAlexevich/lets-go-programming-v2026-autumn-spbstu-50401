package main

import "fmt"

func main() {
	departments := 0
	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Error reading departments count")
		return
	}

	for i := 0; i < departments; i++ {
		employees := 0
		_, err := fmt.Scan(&employees)
		if err != nil {
			fmt.Println("Error reading employee count")
			return
		}
	}

}
