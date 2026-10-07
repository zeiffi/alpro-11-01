package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	hasil := n % 2 == 0

	fmt.Println(hasil)
}