package main

import "fmt"

func main() {
	var n, i int
	var jumlah int = 0

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	i = 1
	for i <= n {
		jumlah += i
		i++
	}

	fmt.Println("Keluaran:", jumlah)
}