package main

import "fmt"

func main() {
	var nilaiUang uint64

	fmt.Print("Masukkan nilai uang	: ")
	fmt.Scanln(&nilaiUang)

	sepuluhRibu := nilaiUang / 10000
	sisa := nilaiUang % 10000

	limaRibu := sisa / 5000
	sisa = sisa % 5000

	seribu := sisa / 1000

	fmt.Println("================ OUTPUT ==================")

	fmt.Println("Nilai cacah uang	:", sepuluhRibu, "lembar 10.000")
	fmt.Println("Nilai cacah uang	:", limaRibu, "lembar 5.000")
	fmt.Println("Nilai cacah uang	:", seribu, "lembar 1.000")

	fmt.Println("==========================================")
}