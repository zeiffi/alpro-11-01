package main

import "fmt"

func main() {
	var jenis string
	var jam int
	fmt.Scan(&jenis, &jam)

	var total int
	if jenis == "Mobil" {
		total = 5000 + (jam-1)*3000
		if jam > 24 {
			total += 50000
		}
	} else if jenis == "Motor" {
		total = 3000 + (jam-1)*2000
		if jam > 24 {
			total += 30000
		}
	}

	fmt.Println(total)
}