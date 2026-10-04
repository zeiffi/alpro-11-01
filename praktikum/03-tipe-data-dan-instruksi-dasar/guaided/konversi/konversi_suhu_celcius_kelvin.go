package main

import "fmt"

func main() {
	var celcius float64;

	fmt.Print("masukan duhu	: ")
	fmt.Scan(&celcius)

	fmt.Println(celcius + 273)
}