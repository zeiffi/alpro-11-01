package main

import "fmt"

func main() {
	var n, lebar, tersempit int

	fmt.Println("Masukkan jumlah daun:")
	fmt.Scan(&n)
	fmt.Println("Masukkan lebar daun:")
	for i := 0; i < n; i++ {
		fmt.Scan(&lebar)
		if i == 0 || lebar < tersempit {
			tersempit = lebar
		}
}
fmt.Println("Lebar daun yang paling sempit:", tersempit)
}