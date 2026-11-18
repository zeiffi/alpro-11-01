# <h1 align="center">Tugas Pendahuluan Modul 10 - FUNGSI</h1>
<p align="center">MUHAMMAD SHEIF FAHIR - 109092600016</p>

### 1. raport.go

```go
package main

import "fmt"

const NMAX int = 100

type student struct {
	nim, name string
	grade     float64
}

type students [NMAX]student

func cariNIM(T students, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].nim == nim {
			return i
		}
	}
	return -1
}

func rataRata(T students, n int) float64 {
	var total float64 = 0.0
	for i := 0; i < n; i++ {
		total += T[i].grade
	}
	return total / float64(n)
}

func main() {
	var T students
	var n int = 0
	var nim, name string
	var grade float64

	for {
		fmt.Scan(&nim)
		if nim == "STOP" {
			break
		}
		fmt.Scan(&name, &grade)

		if cariNIM(T, n, nim) == -1 && n < NMAX {
			T[n].nim = nim
			T[n].name = name
			T[n].grade = grade
			n++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(T[i].nim, T[i].name, T[i].grade)
	}

	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T, n))
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/sisa/output.png)


#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

### 2. [nama_soal, misal: kalkulator.go]

```go
[Tempelkan kode program di sini, contoh: kalkulator.go]
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/konversi/output.png)


#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

## Kesimpulan
[Tuliskan kesimpulan yang menjawab tujuan praktikum berdasarkan hasil yang diperoleh.]