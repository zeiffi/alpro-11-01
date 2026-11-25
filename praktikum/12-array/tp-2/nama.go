package main

import "fmt"

const NMAX int = 9999

type arrString [NMAX]string

func bacaData(A *arrString, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A arrString, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i])
		if i < n-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(".")
}

func sisipData(A *arrString, n *int, p int, x string) {
	// Geser elemen ke kanan dari indeks terakhir sampai p-1
	for i := *n; i >= p; i-- {
		A[i] = A[i-1]
	}
	A[p-1] = x
	*n++
}

func hapusData(A *arrString, n *int, q int) {
	// Geser elemen ke kiri menimpa indeks q-1
	for i := q - 1; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func main() {
	var data arrString
	var n, p, q int
	var namaBaru string

	fmt.Scan(&n)
	bacaData(&data, n)
	fmt.Scan(&p, &namaBaru)
	fmt.Scan(&q)

	fmt.Print("Data awal: ")
	cetakData(data, n)

	sisipData(&data, &n, p, namaBaru)
	fmt.Print("Setelah penyisipan: ")
	cetakData(data, n)

	hapusData(&data, &n, q)
	fmt.Print("Setelah penghapusan: ")
	cetakData(data, n)
}
