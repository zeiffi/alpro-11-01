package main

import "fmt"

func main() {
	var x, y bool

	fmt.Print("masukan nilai ke1	: ")
	fmt.Scan(&x)
	fmt.Print("masukan nilai ke2	: ")
	fmt.Scan(&y)

	fmt.Print("hasil ke1	: ")
	fmt.Println(x)
	fmt.Print("hasil ke	: ")
	fmt.Println(y)
}