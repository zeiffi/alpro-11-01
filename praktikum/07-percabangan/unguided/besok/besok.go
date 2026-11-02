package main

import "fmt"

func isKabisat(tahun int) bool {
	return (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)
}

func getJumlahHari(bulan, tahun int) int {
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isKabisat(tahun) {
			return 29
		}
		return 28
	default:
		return 0
	}
}

func main() {
	var t, b, th int
	fmt.Scan(&t, &b, &th)

	maxHari := getJumlahHari(b, th)

	if b < 1 || b > 12 || t < 1 || t > maxHari {
		fmt.Println("Tidak valid")
		return
	}

	fmt.Println("Valid")

	t++
	if t > maxHari {
		t = 1
		b++
		if b > 12 {
			b = 1
			th++
		}
	}

	fmt.Printf("%d/%d/%d\n", t, b, th)
}