package main

import (
	"fmt"
	"time"

	"github.com/JholmanDanielReinaToledo/CalculatorConsole/functions"
)

func main() {
	continueBucle := true
	var option int

	for continueBucle {
		fmt.Print("\033[H\033[2J") // Clean console
		option = 0                 // reset the option
		fmt.Println("_______MENU_______")
		fmt.Println("Choose your option typing the number")
		fmt.Println("1. Sum 2 numbers")
		fmt.Println("2. Subtraction 2 numbers")
		fmt.Println("3. multiply 2 numbers")
		fmt.Println("4. Divide 2 numbers")
		fmt.Println("99. Exit")

		fmt.Scanln(&option) // read input from user

		switch option {
		case 1:
			functions.GenericOperation(functions.Sum)
			time.Sleep(3 * time.Second)

		case 2:
			functions.GenericOperation(functions.Subtraction)
			time.Sleep(3 * time.Second)

		case 3:
			functions.GenericOperation(functions.Multiply)
			time.Sleep(3 * time.Second)

		case 4:
			functions.GenericOperation(functions.Divide)
			time.Sleep(3 * time.Second)

		case 99:
			continueBucle = false

		default:
			fmt.Println("Unknown option")
		}
	}

	fmt.Println("BYEEEE")
}
