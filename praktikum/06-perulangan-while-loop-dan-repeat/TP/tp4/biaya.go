package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan tahun ke-n: ")
	fmt.Scan(&n)

	var biaya float64 = 10000000
	var tahun int = 1

	for tahun < n {
		biaya = biaya + (biaya * 0.06)
		tahun++
	}

	fmt.Printf("Biaya kuliah pada tahun ke-%d: %.0f\n", n, biaya)
}