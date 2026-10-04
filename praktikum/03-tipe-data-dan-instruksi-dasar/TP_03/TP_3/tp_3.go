package main

import "fmt"

func main (){
	var x float64
	var y float64 = 1.6

	fmt.Print("masukan mil :")
	fmt.Scanln(&x)

	y = x * y
	
	fmt.Print("hasil km :")
	fmt.Printf( "%.1f", y)
}