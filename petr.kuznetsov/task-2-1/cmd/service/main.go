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
		minTemp := 15
		maxTemp := 30
		for j := 0; j < employees; j++ {
			locTemp := 0
			borderSign := ' '
			_, err := fmt.Scanf("%c= %d", &borderSign, &locTemp)
			if err != nil {
				fmt.Println("Error reading temperature")
				return
			}

			switch borderSign {
			case '>':
				minTemp = max(minTemp, locTemp)
			case '<':
				maxTemp = min(maxTemp, locTemp)
			default:
				fmt.Println("Incorrect format")
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}

		}

	}

}
