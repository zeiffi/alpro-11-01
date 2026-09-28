package main

import "fmt"

func main() {
	var celcius float64

	//membaca input
	fmt.Print("masukan celcius	: ")
	fmt.Scanln(&celcius)

	//menghitung total skor, rata rata, dan menampilkan output
	fmt.Println("================= OURPUT =================")

	fmt.Println("suhu dalam reamur	: ", celcius*4/5)
	fmt.Println("suhu dalam fahrenheit	: ", celcius*9/5+32)
	fmt.Println("suhu dalam kelvin	:", celcius+273.15)

	fmt.Println("==========================================")
}