package main

import "fmt"

func dayInWeek(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursdat")
	default:
		fmt.Println("Weekend")

	}
}
func main() {
	dayInWeek(5)
}
