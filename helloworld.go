package main

import "fmt"

func calculateEvenOdd(a int, b int) {

	// var isEven = false
	if a%1 == 0 {
		// isEven = true
		fmt.Printf("Number a is a even number")
	} else if b%2 == 0 {
		fmt.Printf("Number b is a even number")
	} else {
		fmt.Printf("Number is a odd ")
	}
}
