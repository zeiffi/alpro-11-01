package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah siswa: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Jumlah siswa minimal 2 untuk mencari dua nilai terendah.")
		return
	}

	var nama, namaMin1, namaMin2 string
	var nilai, min1, min2 int

	fmt.Print("Masukkan nama dan nilai: ")
	fmt.Scan(&nama, &nilai)
	min1 = nilai
	namaMin1 = nama

	fmt.Print("Masukkan nama dan nilai: ")
	fmt.Scan(&nama, &nilai)
	if nilai < min1 {
		min2 = min1
		namaMin2 = namaMin1
		min1 = nilai
		namaMin1 = nama
	} else {
		min2 = nilai
		namaMin2 = nama
	}

	i := 3
	for i <= n {
		fmt.Print("Masukkan nama dan nilai: ")
		fmt.Scan(&nama, &nilai)

		if nilai < min1 {
			min2 = min1
			namaMin2 = namaMin1
			min1 = nilai
			namaMin1 = nama
		} else if nilai < min2 {
			min2 = nilai
			namaMin2 = nama
		}

		i++
	}

	fmt.Println("Siswa dengan nilai terendah:", namaMin1)
	fmt.Println("Siswa dengan nilai kedua terendah:", namaMin2)
}