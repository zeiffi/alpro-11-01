package main

import "fmt"

func main() {
	var celcius float64

	//membaca input
	fmt.Print("masukan celcius  : ")
	fmt.Scanln(&celcius)

	//menghitung suhu dan menampilkan output
	fmt.Println("================= OUTPUT =================")

	fmt.Println("suhu dalam reamur :", celcius*4/5)
	fmt.Println("suhu dalam fahrenheit :", celcius*9/5+32)
	fmt.Println("suhu dalam kelvin :", celcius+273.15)

	fmt.Println("==========================================")
}