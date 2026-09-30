package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("masukan jumlah kue	:")
	fmt.Scanln(&x)

	fmt.Print("masukan jumlah kue	:")
	fmt.Scanln(&y)

	fmt.Println("hasil kue sisa bagi	:", x % y)
}