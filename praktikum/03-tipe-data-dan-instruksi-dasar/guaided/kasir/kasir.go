package main

import "fmt"

func main() {
	var x int

	fmt.Print("masukan nominal : ")
	fmt.Scan(&x)

	var sepuluhRirbu int = x / 10000
	var sisa int = x % 10000

	var limaRibuan int = sisa / 5000
	 sisa  = sisa % 5000

	var seRibuan int = sisa / 1000
	
	fmt.Println(sepuluhRirbu, limaRibuan, seRibuan)
}
