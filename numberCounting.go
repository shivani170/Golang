package main

import "fmt"

func main() {
	var inputNumber int
	fmt.Scanln(&inputNumber)

	fmt.Scanf("%d", &inputNumber)
	count := 0
	for inputNumber > 0 {
		inputNumber = inputNumber / 10
		count++
	}

	// for count := 0; count < inputNumber.length; count++{
	// 	fmt.Println(count)
	// }

	fmt.Printf("The number of digits in the given number is: %d", count)

}
