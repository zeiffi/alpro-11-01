package main

import "fmt"

func main() {
	var hari int

	fmt.Print("masukan jumlah hari : ")
	fmt.Scan(&hari)

	// Menggunakan 360 hari untuk 1 tahun (tahun komersial)
	var tahun int = hari / 360
	var sisaHari int = hari % 360

	var bulan int = sisaHari / 30
	sisaHari = sisaHari % 30

	var minggu int = sisaHari / 7
	sisaHari = sisaHari % 7

	fmt.Println("================ HASIL HARI =================")

	fmt.Println("tahun :", tahun)
	fmt.Println("bulan :", bulan)
	fmt.Println("minggu :", minggu)
	fmt.Println("hari :", sisaHari)

	fmt.Println("=============================================")
}