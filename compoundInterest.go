package main

import (
	"fmt"
	"math"
)

func compoundInterest(principle float64, interest float64, time float64) {
	var CI float64
	var ciFuture = principle * (math.Pow((1 + interest/100), time))
	CI = ciFuture - principle

	fmt.Println(CI)
	fmt.Println("The Future Compound Interest  = ", ciFuture)
}

func main() {
	compoundInterest(500, 20, 2)
}
