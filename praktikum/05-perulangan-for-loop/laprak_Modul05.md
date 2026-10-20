# <h1 align="center">Laporan Praktikum Modul [Nomor Modul] - [Judul Modul/Topik]</h1>
<p align="center">[Nama Praktikan] - [NIM]</p>

## Dasar Teori

### A. Konsep Dasar Struktur Perulangan (Looping) dan For Loop di Go
Menurut konsep Algoritma dan Pemrograman Telkom University Purwokerto, perulangan (looping) adalah struktur kontrol yang digunakan untuk mengeksekusi satu blok kode secara berulang-ulang selama suatu kondisi tertentu masih terpenuhi. Berbeda dengan bahasa pemrograman lain yang menyediakan banyak kata kunci perulangan (for, while, do-while), bahasa pemrograman Go hanya menyediakan satu kata kunci utama yaitu for. Meskipun hanya menggunakan satu kata kunci, fleksibilitas bentuk sintaksnya mampu menggantikan seluruh fungsi perulangan tersebut, mulai dari bentuk klasik tiga komponen, pengganti while, hingga perulangan tanpa henti (infinite loop).

### B. Bentuk Pengulangan dan Perulangan Bersarang (Nested Loop) dalam Go

#### 1. For dengan Tiga Konponen (ala C):
Bentuk klasik menggunakan inisialisasi, kondisi, dan post-statement
```go
for i := 1; i <= 5; i++
```

#### 2. For sebagai Pengganti While :
Hanya menyertakan kondisi saja tanpa inisialisasi dan post-statement
```go
for counter < 5
```

#### 3. For Tanpa Kondisi (Infinite Loop):
Perulangan tanpa henti yang biasanya dikendalikan menggunakan pernyataan break di dalam blok kodenya
```go 
for { ... } i++
```

#### 4. For-Range:
Digunakan khusus untuk melakukan iterasi atau menelusuri elemen pada koleksi data seperti slice, array, string, atau map
```go 
for index, value := range buah
```

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. tabungan.go
```go
package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scan(&a, &b, &n)
	
	total := 0
	tabunganHariIni := a

	for i := 1; i <= n; i++ {
		total = total + tabunganHariIni
		tabunganHariIni = tabunganHariIni + b
		
	}
	fmt.Println(total)
}
```
#### Deskripsi
Program ini menghitung total tabungan Budi selama $n$ hari dengan menjumlahkan setoran awal $a$ yang terus bertambah sebesar $b$ setiap harinya menggunakan perulangan for

### 2. tebesar.go

```go
package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	var angka, max int
	fmt.Scan(&angka)
	
	for i := 2; i < n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}
	}
	fmt.Println("================== nilai max ==================")
	fmt.Println(max)
	fmt.Println("===============================================")
}
```
#### Deskripsi
 Program ini bertugas mencari nilai atau angka paling tinggi dari sekumpulan $n$ bilangan bulat yang dimasukkan oleh pengguna dengan cara membandingkannya satu persatu.


### 3. bisokop.go

```go
package main

import "fmt"

func main() {
	var t int
	fmt.Print("input t (jumlah transaksi): ")
	fmt.Scan(&t)

	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisFilm, hari string
		var jumlahTiket int

		fmt.Printf("input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0

		if jenisFilm == "Reguler" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 40000
			} else {
				hargaSatuan = 50000
			}
		} else if jenisFilm == "3D" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 60000
			} else {
				hargaSatuan = 70000
			}
		} else if jenisFilm == "IMAX" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 80000
			} else {
				hargaSatuan = 90000
			}
		}

		totalHarga := hargaSatuan * jumlahTiket

		if jumlahTiket >= 5 {
			diskon := (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}

		fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalHarga)

		totalPendapatan = totalPendapatan + totalHarga
	}

	fmt.Println("output total pendapatan:", totalPendapatan)
}
```
#### Deskripsi
Program ini menghitung total harga tiket bioskop dari beberapa transaksi berdasarkan jenis film dan hari, lalu otomatis memberikan diskon 10% jika pembelian tiket dalam satu transaksi mencapai 5 atau lebih.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. prima.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor dari %d: ", n)
	jumlahFaktor := 0
	
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
			jumlahFaktor++
		}
	}
	fmt.Printf("(%d faktor)\n", jumlahFaktor)

	if jumlahFaktor == 2 {
		fmt.Printf("%d adalah bilangan prima\n", n)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", n)
	}
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/unguided/perima.go/output.png)


#### Deskripsi
Program ini mencari seluruh faktor pembagi dari sebuah bilangan bulat positif $n$ dan menghitung jumlahnya untuk menentukan apakah bilangan tersebut termasuk bilangan prima atau bukan
### 2. acct.go

```go
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
```

##### Output
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/unguided/acct/output.png)

#### Deskripsi
Program ini mengevaluasi rata-rata dari 6 nilai tes ECCT untuk setiap mahasiswa baru guna menentukan kelulusan sertifikat (minimal 3.33) sekaligus menghitung total mahasiswa yang lulus.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Struktur perulangan for dalam bahasa Go sangat fleksibel—mulai dari bentuk klasik tiga komponen, pengganti while, hingga for-range—yang diterapkan pada tugas terbimbing seperti menghitung tabungan, mencari nilai terbesar, dan rekapitulasi bioskop, serta tugas mandiri seperti mencari faktor prima dan mengevaluasi rata-rata nilai sertifikat mahasiswa.
1. Hidayat, R., & Pratama, A. (2025)