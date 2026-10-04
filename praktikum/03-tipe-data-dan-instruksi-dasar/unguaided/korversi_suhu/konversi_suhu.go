package main

import "fmt"

func main() {
	var celcius int
	var remur int

	fmt.Print("masukan suhu dalam celcius :")
	fmt.Scan(&celcius)

	remur = (celcius * 4) / 5

	fmt.Println("================ HASIL KONVERSI =================")

	fmt.Print("hasil dari konversi remur : ")
	fmt.Println(remur)

	fmt.Println("=================================================")
}