package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	var angka int
	max := 0
	fmt.Scan(&angka)

	for i := 2; i < n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}
	fmt.Println("================== nilai max ==================")
	fmt.Println(max)
	fmt.Println("===============================================")
}
