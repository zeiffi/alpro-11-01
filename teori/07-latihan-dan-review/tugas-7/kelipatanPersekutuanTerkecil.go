package main

import "fmt"

func main() {
	var n, m int

	fmt.Print("Masukkan dua bilangan bulat positif: ")
	fmt.Scan(&n, &m)

	a, b := n, m // a = kelipatan N saat ini, b = kelipatan M saat ini
	for a != b {
		if a < b {
			a += n
		} else {
			b += m
		}
	}
	fmt.Println("hasil persekutuan terkecil:", a)
}