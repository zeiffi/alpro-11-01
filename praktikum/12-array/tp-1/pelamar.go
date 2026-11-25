package main

import "fmt"

const NMAX int = 100

type applicant struct {
	id, name     string
	testGrade    float64
	testDuration int
}

type applicants [NMAX]applicant

func salinLolos(A applicants, n int, batas float64, B *applicants, m *int) {
	*m = 0
	for i := 0; i < n; i++ {
		if A[i].testGrade >= batas {
			B[*m] = A[i]
			*m++
		}
	}
}

func rataDurasi(B applicants, m int) float64 {
	var total int = 0
	for i := 0; i < m; i++ {
		total += B[i].testDuration
	}
	return float64(total) / float64(m)
}

func main() {
	var data, lolos applicants
	var n, m int
	var batas float64
	var tempID string

	n = 0
	fmt.Scan(&tempID)
	for tempID != "END" && n < NMAX {
		data[n].id = tempID
		fmt.Scan(&data[n].name, &data[n].testGrade, &data[n].testDuration)
		n++
		fmt.Scan(&tempID)
	}
	fmt.Scan(&batas)

	salinLolos(data, n, batas, &lolos, &m)

	fmt.Println("Data awal:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %.1f %d\n", data[i].id, data[i].name, data[i].testGrade, data[i].testDuration)
	}

	fmt.Printf("Pelamar lolos (nilai >= %.1f):\n", batas)
	if m == 0 {
		fmt.Println("Tidak ada pelamar yang lolos")
	} else {
		for i := 0; i < m; i++ {
			fmt.Printf("%s %s %.1f %d\n", lolos[i].id, lolos[i].name, lolos[i].testGrade, lolos[i].testDuration)
		}
		fmt.Printf("Jumlah pelamar lolos: %d\n", m)
		fmt.Printf("Rata-rata durasi: %.2f menit\n", rataDurasi(lolos, m))
	}
}
