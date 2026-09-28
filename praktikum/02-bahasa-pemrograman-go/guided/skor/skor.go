package main

import "fmt"

func main() {
	var nama string
	var skorMatematika, skorBahasaInggris int

	//membaca input
	fmt.Print("masukan nama siswa	:")
	fmt.Scan(&nama)
	fmt.Print("masukan skor math	:")
	fmt.Scan(&skorMatematika)
	fmt.Print("masukan skor inggris	:")
	fmt.Scan(&skorBahasaInggris)

	//menghitung total & rata rata (pembagian bilangan bulat)
	total  := skorMatematika + skorBahasaInggris
	rataRata := total/2

	//menampilkan output
	fmt.Println("=============== DATA ===============") 

	fmt.Println("nama siswa	:", nama)
	fmt.Println("total		:", total)
	fmt.Println("rata rata	:", rataRata)
	
	fmt.Println("====================================")
}