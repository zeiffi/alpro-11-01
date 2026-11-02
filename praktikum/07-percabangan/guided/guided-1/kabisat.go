package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("masukan tahun : ")
	fmt.Scan(&tahun)

	if (tahun % 400 == 0 && tahun % 100 != 0) || (tahun % 400 != 0) {
		fmt.Print("iya")
	} else {
		fmt.Print("tidak")
	}
}

