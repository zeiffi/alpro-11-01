package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Masukkan nilai a	: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai b	: ")
	fmt.Scanln(&b)

	fmt.Println("================ OUTPUT ==================")

	fmt.Println("Hasil penjumlahan	:", a+b)
	fmt.Println("Hasil pengurangan	:", a-b)
	fmt.Println("Hasil perkalian	:", a*b)
	fmt.Println("Hasil pembagian	:", a/b)
	fmt.Println("Hasil modulo		:", math.Mod(a, b))

	fmt.Println("==========================================")
}