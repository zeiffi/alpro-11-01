package main

import "fmt"

func main() {
	var jari float64
	var pi float64 = 3.14

	//membaca input
	fmt.Print("masukan jari-jari : ")
	fmt.Scanln(&jari)

	//menghitung luas lingkaran dan menampilkan output
	fmt.Println("========================= OUTPUT =========================")

	fmt.Println("luas lingkaran :", pi * jari * jari)
	
	fmt.Println("==========================================================")
}