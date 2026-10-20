package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor dari %d: ", n)
	jumlahFaktor := 0
	
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
			jumlahFaktor++
		}
	}
	fmt.Printf("(%d faktor)\n", jumlahFaktor)

	if jumlahFaktor == 2 {
		fmt.Printf("%d adalah bilangan prima\n", n)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", n)
	}
}
