package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nilai)

	fmt.Println("========================================")
	if nilai <= 50 {
		fmt.Println("Nilai anda (E)")
	}else if nilai > 50 && nilai <= 70 {
		fmt.Println("Nilai anda (C)")
	}else if nilai > 70 && nilai <= 85 {
		fmt.Println("Nilai anda (B)")
	}else if nilai > 85 && nilai <= 100 {
		fmt.Println("Nilai anda (A)")
	}else {
		fmt.Println("Nilai tidak valid")
	}
	fmt.Println("========================================")
}