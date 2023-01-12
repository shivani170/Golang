package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	v1 := m["k1"]
	fmt.Println("Value in amp", v1)
	m["k2"] = 8

	fmt.Println("Value Before delete in amp", m)

	delete(m, "k2")
	fmt.Println("Value Before delete in amp", m)

	_, prs1 := m["k1"]
	fmt.Println("Catching the value in ", prs1)

	_, prs2 := m["k2"]
	fmt.Println("Catching the value in ", prs2)
}
