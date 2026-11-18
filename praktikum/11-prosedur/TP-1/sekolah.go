package main

import "fmt"

const NMAX int = 1001

type tabMurid [NMAX]string

func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	*ketemu = false
	for i := 0; i < n; i++ {
		if T[i] == nama {
			*ketemu = true
			*posisi = i + 1 // urutan absen dimulai dari 1
			return
		}
	}
}

func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	*jumlah = 0
	for i := 0; i < n; i++ {
		if len(T[i]) > 0 && T[i][0] == h {
			*jumlah++
		}
	}
}

func main() {
	var T tabMurid
	var n, posisi, jumlah int
	var namaCari string
	var ketemu bool

	// Input jumlah murid
	fmt.Scan(&n)

	// Input daftar nama murid
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	// Input nama yang dicari
	fmt.Scan(&namaCari)

	// Pemanggilan prosedur
	cariMurid(T, n, namaCari, &posisi, &ketemu)
	hitungHuruf(T, n, namaCari[0], &jumlah)

	// Output hasil
	if ketemu {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}
	fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", namaCari[0], jumlah)
}