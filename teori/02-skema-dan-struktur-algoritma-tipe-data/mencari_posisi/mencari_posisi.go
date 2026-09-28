package main

import "fmt"

func main() {
	var posisi, posisi0, kecepatan, waktu int

	fmt.scan(&posisi0, &kecepatan, &waktu)

	posisi = posisi0 + (kecepatan*waktu)

	fmt.Println("posisi atau jarak adalah : ", posisi)
}