package main

import "fmt"

func main() {
	var p, q int

	fmt.Scan(&p, &q)

	OR := (p % 2 == 0) || (q % 2 == 0)
	AND := (p % 2 != 0) && (q % 2 != 0)
	NOT :=  !(p == q)

	fmt.Println(OR, AND, NOT)
}