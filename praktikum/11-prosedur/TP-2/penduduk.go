package main

import "fmt"

const NMAX int = 1000

type penduduk struct {
	nama  string
	tahun int
	kota  string
}

type tabPenduduk [NMAX]penduduk

func cariTahun(T tabPenduduk, n, th int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i].tahun == th {
			*idx = i
			return // berhenti di data pertama yang cocok
		}
	}
}

func tertuaTermuda(T tabPenduduk, n int, tua, muda *penduduk) {
	*tua = T[0]
	*muda = T[0]

	for i := 1; i < n; i++ {
		if T[i].tahun < (*tua).tahun {
			*tua = T[i]
		}
		if T[i].tahun > (*muda).tahun {
			*muda = T[i]
		}
	}
}

func main() {
	var T tabPenduduk
	var n, thCari, idx int
	var tua, muda penduduk

	// Input jumlah data penduduk
	fmt.Scan(&n)

	// Input data nama, tahun lahir, dan kota
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].kota)
	}

	// Input tahun lahir yang dicari
	fmt.Scan(&thCari)

	// Pemanggilan prosedur
	cariTahun(T, n, thCari, &idx)
	tertuaTermuda(T, n, &tua, &muda)

	// Output hasil
	if idx != -1 {
		fmt.Printf("%s %d %s ditemukan di index ke-%d\n", T[idx].nama, T[idx].tahun, T[idx].kota, idx)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
	fmt.Printf("Penduduk tertua: %s (%d)\n", tua.nama, tua.tahun)
	fmt.Printf("Penduduk termuda: %s (%d)\n", muda.nama, muda.tahun)
}