package main

import "fmt"

func main() {
	var m int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&m)

	jumlahLolos := 0

	for i := 1; i <= m; i++ {
		var p, v, ga, fs, ci, tc float64
		
		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		rataRata := (p + v + ga + fs + ci + tc) / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Mendapat Sertifikat")
			jumlahLolos++
		} else {
			fmt.Println("Tidak Mendapat Sertifikat")
		}
	}
	fmt.Printf("Jumlah yang mendapat sertifikat: %d\n", jumlahLolos)
}