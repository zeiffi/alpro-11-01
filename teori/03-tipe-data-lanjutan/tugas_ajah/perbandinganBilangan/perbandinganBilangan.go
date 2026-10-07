package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	lebihBesar := x > y
	samaDengan := x == y
	lebihKecil := x < y

	fmt.Println(lebihBesar, samaDengan, lebihKecil)
}