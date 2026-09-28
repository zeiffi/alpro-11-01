package main

import "fmt"

func main() {
	var jari float64
	var pi float64 = 3.14

	//membaca input
	fmt.Print("masukan jari-jari : ")
	fmt.Scanln(&jari)

	//menghitung tonton skor, rata-rata, dan menampilkan input
	fmt.Println("========================= OUTPUT =========================")

	fmt.Println("luas lingkaran	:", pi*loat64(jari)*float64(jari))

	fmt.Println("==========================================================")
}