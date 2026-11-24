package main

import "fmt"

const NMAX int = 999

type gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

type tabGempa [NMAX]gempa

func kategori(m float64) string {
	if m < 4.0 {
		return "Minor"
	} else if m < 6.0 {
		return "Sedang"
	} else if m < 7.0 {
		return "Kuat"
	} else {
		return "Besar"
	}
}

func rataMagnitudo(T tabGempa, n int) float64 {
	var total float64 = 0.0
	for i := 0; i < n; i++ {
		total += T[i].magnitudo
	}
	return total / float64(n)
}

func idxTerkuat(T tabGempa, n int) int {
	var maxIdx int = 0
	for i := 1; i < n; i++ {
		// Menggunakan '>' agar gempa pertama yang muncul tetap terpilih jika magnitudonya sama
		if T[i].magnitudo > T[maxIdx].magnitudo {
			maxIdx = i
		}
	}
	return maxIdx
}

func hitungKategori(T tabGempa, n int, k string) int {
	var count int = 0
	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			count++
		}
	}
	return count
}

func main() {
	var data tabGempa
	var n int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].nama, &data[i].tahun, &data[i].magnitudo)
	}

	for i := 0; i < n; i++ {
		kat := kategori(data[i].magnitudo)
		fmt.Printf("%s %d %.1f %s\n", data[i].nama, data[i].tahun, data[i].magnitudo, kat)
	}

	avg := rataMagnitudo(data, n)
	strongestIdx := idxTerkuat(data, n)
	totalBesar := hitungKategori(data, n, "Besar")

	fmt.Printf("Rata-rata magnitudo: %.2f\n", avg)
	fmt.Printf("Gempa terkuat: %s %d %.1f\n", data[strongestIdx].nama, data[strongestIdx].tahun, data[strongestIdx].magnitudo)
	fmt.Printf("Jumlah gempa kategori Besar: %d\n", totalBesar)
}