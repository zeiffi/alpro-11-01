package main

import "fmt"

func main() {
	var lampu_satatus string 

	fmt.Print("Masukkan warna lampu (merah/kuning/hijau): ")
	fmt.Scanln(&lampu_satatus)

	fmt.Println("========================================")
	if lampu_satatus == "merah" {
		fmt.Println("Lampu merah, berhenti!")
	} else if lampu_satatus == "kuning" {
		fmt.Println("Lampu kuning, bersiap-siap!")
	} else if lampu_satatus == "hijau" {
		fmt.Println("Lampu hijau, jalan!")
	} else {
		fmt.Println("tidak ada daftar lampu")
	}
	fmt.Println("========================================")
}