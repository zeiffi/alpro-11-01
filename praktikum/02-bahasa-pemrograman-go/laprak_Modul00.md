# <h1 align="center">Laporan Praktikum Modul 1 - dasar pemograman go</h1>
<p align="center">muhammad seheif fahri - 109092600016</p>

## Dasar Teori

### A. pengenalan dasar bahasa pemograman go
Go atau yang sering disebut Golang adalah bahasa pemrograman prosedural yang dikembangkan di Google oleh Robert Griesemer, Rob Pike, dan Ken Thompson pada tahun 2007. Bahasa ini kemudian dirilis untuk publik sebagai perangkat lunak open-source pada tahun 2009.

Popularitas Golang mulai menanjak sejak diadopsi untuk membangun Docker pada tahun 2011. Saat ini, Go sangat diandalkan dalam pengembangan Backend API dengan arsitektur microservices. Bahkan, banyak teknologi modern kini beralih menggunakan Go daripada bahasa C, seperti Kubernetes, Prometheus, CockroachDB, dan lain-lain.

(Sumber: Irvan Eksa Mahendra, 2021).


### B.Struktur Program dan Package dalam Go (Golang)


#### 1.  Memahami Peran package main dan func main()
Aturan Package: Setiap file program dalam Go wajib mendeklarasikan sebuah package. Dalam sebuah proyek, setidaknya harus ada satu file yang menggunakan package main.

Eksekusi Utama: File dengan package main merupakan bagian yang pertama kali dijalankan oleh program.

Fungsi main(): Di dalam package main, wajib terdapat sebuah fungsi utama bernama main(). Fungsi inilah yang menjadi titik awal eksekusi utama saat program dijalankan.

(Sumber: Noval Agung Prayogo, 2019)



#### 2. Deklarasi Variabel dan Tipe Data di Go
Go adalah bahasa statically typed, artinya tipe data variabel dicek saat kompilasi. Namun, Go memiliki fitur canggih bernama Type Inference.

Deklarasi Variabel

Ada dua cara umum untuk membuat variabel:

Cara 1: Deklarasi Eksplisit

Digunakan ketika Anda ingin menentukan tipe data secara manual atau mendeklarasikan tanpa nilai awal.
```go
var nama string = "mitha"
var umur int = 17
```
Cara 2: Short Variable Declaration (:=)

Ini adalah cara paling umum di dalam fungsi. Go akan otomatis menebak tipe datanya.

```go
negara := "Indonesia" // Go tahu ini String
skor := 55,3          // Go tahu ini Float64
```
 

## Guided

### 1. skor.go

```go
package main

import "fmt"

func main() {
	var nama string
	var skorMatematika, skorBahasaInggris int

	//membaca input
	fmt.Print("masukan nama siswa	:")
	fmt.Scan(&nama)
	fmt.Print("masukan skor math	:")
	fmt.Scan(&skorMatematika)
	fmt.Print("masukan skor inggris	:")
	fmt.Scan(&skorBahasaInggris)

	//menghitung total & rata rata (pembagian bilangan bulat)
	total  := skorMatematika + skorBahasaInggris
	rataRata := total/2

	//menampilkan output
	fmt.Println("=============== DATA ===============") 

	fmt.Println("nama siswa	:", nama)
	fmt.Println("total		:", total)
	fmt.Println("rata rata	:", rataRata)
	
	fmt.Println("====================================")
}

```
#### Deskripsi
Program di atas membaca input nama siswa dengan variable (string), skorMtk (uint8), dan skorBhsInggris (uint8). Kemudian program akan menghitung total dari kedua nilai variable (skorMtk, skroBhsInggris) untuk menghitung seluruh total nilai, setelah itu total nilai akan masuk ke proses pembagian ((skorMtk+skorBhsInggris)/2) guna menghitung rata - rata dari nilai atau skor siswa yang telah diinputkan.

catatan: bisa menggunakan tipe data (uint8) seperti di atas atau menggunakan (int) saja dan bisa membuat variable "totalSkor" untuk menyimpan nilai dari variable skorMtk + skorBhsInggris begitu juga dengan skor rata - rata kalian bisa menggunakan/membuat variable sendiri guna menyimpan nilai nilai rata - rata skor siswa.

### 2. tukar.go

```go
package main

import "fmt"

func main() {
	var a,b int

	//membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//menukar a dan b
	a,b=b,a

	//output
	fmt.Println(a)
	fmt.Println(b)
}
```
#### Deskripsi
Program di atas membaca nilai yang diinput ke dalam variable a (int) dan b (int) lalu menampilkan hasil inputan kita secara tertukar contoh: kita menginputkan a = 20, b = 30 hasilnya Nilai a: 30, Nilai b: 20.


### 3. lingkaran.go

```go
package main

import "fmt"

func main() {
	var jari float64
	var pi float64 = 3.14

	//membaca input
	fmt.Print("masukan jari-jari : ")
	fmt.Scanln(&jari)

	//menghitung tonton skor, rata-rata, dan menampilkan input
	fmt.Println("========================= OUTPUT =========================")

	fmt.Println("luas lingkaran	:", pi*loat64(jari)*float64(jari))

	fmt.Println("==========================================================")
}

```
#### Deskripsi
Program di atas menghitung luas lingkaran dengan cara menghitung jari - jari yang diinputkan dan disimpan di variable "r" (float64) dengan "pi (float64). Cara menghitungnya mengikuti rmus menghitung luas lingkaran yaitu: pi _ r _. Dideklarasikan dengan kode:
```go
fmt.Println("Luas lingkaran		:", pi*(r)*(r))
```

### 4. suhu.go

```go
package main

import "fmt"

func main() {
	var celcius float64

	//membaca input
	fmt.Print("masukan celcius  : ")
	fmt.Scanln(&celcius)

	//menghitung suhu dan menampilkan output
	fmt.Println("================= OUTPUT =================")

	fmt.Println("suhu dalam reamur :", celcius*4/5)
	fmt.Println("suhu dalam fahrenheit :", celcius*9/5+32)
	fmt.Println("suhu dalam kelvin :", celcius+273.15)

	fmt.Println("==========================================")
}
```
#### Deskripsi
Program di atas membaca nilai celsius yang kita input lalu dikonvert mennjadi reamur, farenheit, dan kelvin. Nilai yang diinput disimpan dalan variable "celsius" (float64) lalu dikonvert menjadi beberapa jenis suhu.

##### deklarasi

```go
fmt.Println("Suhu dalam Reamur	:", celsius*4/5)
fmt.Println("Suhu dalam Fahrenheit	:", celcius*9/5+32)
fmt.Println("Suhu dalam Kelvin	:", celsius+273.15)
```

## Unguided


### 1. kalkulator

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Masukkan nilai a	: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai b	: ")
	fmt.Scanln(&b)

	fmt.Println("================ OUTPUT ==================")

	fmt.Println("Hasil penjumlahan	:", a+b)
	fmt.Println("Hasil pengurangan	:", a-b)
	fmt.Println("Hasil perkalian	:", a*b)
	fmt.Println("Hasil pembagian	:", a/b)
	fmt.Println("Hasil modulo		:", math.Mod(a, b))

	fmt.Println("==========================================")
}
```

##### Output
!![\[Screenshot Output Unguided\]](unguided/cacahuang/output.png)

#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Praktikum ini bertujuan untuk mengenalkan dasar - dasar bahasa pemrograman Go dan menjelaskan bagaimana cara penggunaannya melalui beberapa latian membuat program yang ada di atas.

### 2. cacahuang

```go
package main

import "fmt"

func main() {
	var nilaiUang uint64

	fmt.Print("Masukkan nilai uang	: ")
	fmt.Scanln(&nilaiUang)

	sepuluhRibu := nilaiUang / 10000
	sisa := nilaiUang % 10000

	limaRibu := sisa / 5000
	sisa = sisa % 5000

	seribu := sisa / 1000

	fmt.Println("================ OUTPUT ==================")

	fmt.Println("Nilai cacah uang	:", sepuluhRibu, "lembar 10.000")
	fmt.Println("Nilai cacah uang	:", limaRibu, "lembar 5.000")
	fmt.Println("Nilai cacah uang	:", seribu, "lembar 1.000")

	fmt.Println("==========================================")
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output.png)


#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

## Referensi
1. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
2. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
