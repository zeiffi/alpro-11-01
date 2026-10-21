# <h1 align="center">Tugas Pendahuluan Modul 05 - PERULANGAN WHILE</h1>
<p align="center">MUHAMMAD SHEIF FAHRI - 109092600016</p>

### 1. lulus.go

```go
package main

import "fmt"

func main() {
	var nilai int

	for {
		fmt.Print("Masukkan nilai Anda: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			fmt.Println("Program selesai.")
			break
		}

		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}
		fmt.Println()
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/TP/tp1/output.png)


#### Deskripsi
Deskripsi: Menggunakan perulangan tak berhingga (for { ... }) untuk meminta input nilai berulang kali. Jika input bernilai -1, perulangan dihentikan dengan break. Jika tidak, program menyeleksi apakah nilai tersebut $\ge 60$ untuk menentukan status kelulusan.

### 2. deret.go
```go
package main

import "fmt"

func main() {
	var n, i int
	var jumlah int = 0

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	i = 1
	for i <= n {
		jumlah += i
		i++
	}

	fmt.Println("Keluaran:", jumlah)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/TP/tp2/output.png)


#### Deskripsi
Deskripsi: Memanfaatkan pola perulangan while (for i <= n) untuk mengakumulasi penjumlahan bilangan dari $1$ hingga $n$. Variabel jumlah akan terus menambahkan nilai i pada setiap iterasi hingga kondisi batas tercapai.

### 3. Dua terendah.go



```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/TP/tp3/output.png)


#### Deskripsi
Deskripsi: Memproses $n$ data nama dan nilai siswa menggunakan perulangan terstruktur. Program melakukan perbandingan bertingkat untuk terus memperbarui dan menyimpan dua nilai serta nama siswa terendah pertama (min1) dan kedua (min2).

### 4. Biaya

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan tahun ke-n: ")
	fmt.Scan(&n)

	var biaya float64 = 10000000
	var tahun int = 1

	for tahun < n {
		biaya = biaya + (biaya * 0.06)
		tahun++
	}

	fmt.Printf("Biaya kuliah pada tahun ke-%d: %.0f\n", n, biaya)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/TP/tp4/output.png)


#### Deskripsi
Deskripsi: Menghitung proyeksi kenaikan biaya kuliah secara iteratif per tahun menggunakan tipe data float64. Setiap tahunnya, biaya dikalikan kenaikan 6% ($1.06$) hingga mencapai tahun ke-$n$, lalu hasilnya dibulatkan menggunakan format %.0f.

## Kesimpulan
Secara keseluruhan, keempat program ini melatih penggunaan pola perulangan while (for kondisi) dalam bahasa Go—yang tidak memiliki keyword while bawaan—untuk menyelesaikan berbagai kasus, mulai dari penghentian proses dengan sentinel (-1), akumulasi deret, pencarian nilai minimum bertingkat, hingga perhitungan matematika berulang (bunga/kenaikan biaya).