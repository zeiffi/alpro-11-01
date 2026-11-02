package main

import(
	"fmt"
	"math"
)

func main() {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Printf("berat melebihi batas maksimum 30kg")
		return 
	}

	var tarifPerKg int

	switch zona {
	case 1:
		tarifPerKg = 10000
	case 2:
		tarifPerKg = 15000
	case 3: 
		tarifPerKg = 25000
	case 4: 
		fmt.Print("zona tidak valid")
		return
	}

	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarifPerKg

	if layanan == "Express" {
		total += total + total/2
		
	}

	fmt.Println(beratBulat)
	fmt.Println(total)
}