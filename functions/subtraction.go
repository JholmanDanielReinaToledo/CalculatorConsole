package functions

import "fmt"

func Subtraction() {
	var num1, num2 float64

	fmt.Println("Input first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Input second number: ")
	fmt.Scanln(&num2)

	result := num1 - num2

	fmt.Printf("The result is: %.2f \n", result)
}
