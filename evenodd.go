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

func main() {

	// fmt.Println("Hello world")

	// Declaring a variable
	// var a int = 10
	// if a%2 == 0 {
	//  fmt.Printf("It is a Printf ")
	//  fmt.Println("It is a Println")
	//  fmt.Print("It is a Print ")
	// }

	// program to check if a number is prime or not
	// const number = 5
	// var isPrime bool = true
	// if number == 1 || number == 0 {
	//  fmt.Println("1 is neither prime nor composite number.") // check if number is equal to 1
	// } else if number > 1 { // check if number is greater than 1
	//  // looping through 2 to number-1
	//  for i := 2; i < number; i++ {
	//    if number%i == 0 {
	//      isPrime = false
	//      break
	//    }
	//  }
	//  if isPrime {
	//    fmt.Println(number, `is a prime number`)
	//  } else {
	//    fmt.Println(number, `${number} is a not prime number`)
	//  }
	// } else {
	//  fmt.Println(number, "The number is not a prime number.")
	// }
	calculateEvenOdd(4, 3)
	// calculateEvenOdd(3)

	fmt.Println("test")
	switch isEven := true; isEven {
	case true:
		fmt.Println("even")
	case false:
		fmt.Println("odd")
	default:
		fmt.Println(isEven)
	}
}
