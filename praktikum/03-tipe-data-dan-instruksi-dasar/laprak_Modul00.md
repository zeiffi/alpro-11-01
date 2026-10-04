# <h1 align="center">Laporan Praktikum Modul [Nomor Modul] - [Judul Modul/Topik]</h1>
<p align="center">[Nama Praktikan] - [NIM]</p>

## Dasar Teori

### A. Konsep Variabel dan Tipe Data (Variables and Data Types)
variabel merupakan lokasi memori bernama yang digunakan untuk menyimpan nilai yang dapat diakses atau diubah selama program berjalan. Tipe data menentukan jenis nilai yang dapat disimpan oleh variabel serta jenis operasi yang dapat dilakukan terhadapnya.
 menurut K.N.king (2008)
  Alan A. Donovan & Brian W. Kernighan (2015)

#### 1. Operator Aritmatika Dasar
Operator aritmatika adalah simbol bawaan dalam bahasa pemrograman yang memanipulasi operan numerik untuk menghasilkan nilai baru (Dale & Weems, 2012). Dalam bahasa Go, operator seperti penjumlahan (+), pengurangan (-), perkalian (*), dan pembagian (/) digunakan untuk menyelesaikan perhitungan matematika langsung, seperti konversi skala suhu dari Celsius ke Kelvin 
```go
K = C + 273
```

#### 2. Operator Modulo dan Algoritma Pembagian Bilangan Bulat
Operator modulo (%) merupakan operator khusus yang mengembalikan sisa hasil bagi dari pembagian dua bilangan bulat (Dale & Weems, 2012). Operator ini sangat penting dalam algoritma pemecahan nilai numerik (decomposition), seperti memecah jumlah uang kembalian ke dalam lembaran pecahan (10.000, 5.000, 1.000) atau mengonversi total jumlah hari menjadi satuan tahun, bulan, minggu, dan sisa hari


<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. konversi_suhu_celcius_kelvin.go

```go
package main

import "fmt"

func main() {
	var celcius float64;

	fmt.Print("masukan duhu	: ")
	fmt.Scan(&celcius)

	fmt.Println(celcius + 273)
}
```
#### Deskripsi
program dalam bahasa Go untuk mengonversi suhu dari derajat Celsius menjadi 
Kelvin dengan menambahkan K = C + 273 

### 2. tukar.go


```go
package main

import "fmt"

func main() {
	var x, y, z int

	fmt.Scan(&x, &y, &z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println(x, y, z)
}
```
#### Deskripsi
 program dalam bahasa Go untuk mempertukarkan nilai bilangan bulat x, y, dan z  

### 3. kasi.go


```go
package main

import "fmt"

func main() {
	var x int

	fmt.Print("masukan nominal : ")
	fmt.Scan(&x)

	var sepuluhRirbu int = x / 10000
	var sisa int = x % 10000

	var limaRibuan int = sisa / 5000
	 sisa  = sisa % 5000

	var seRibuan int = sisa / 1000
	
	fmt.Println(sepuluhRirbu, limaRibuan, seRibuan)
}

```
#### Deskripsi
menghitung lembaran uang 1000, 5000, dan 10000 agar mengetahui berapa lembar yang di dapat dari nominal yang diberikan

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. konversi suhu

```go
package main

import "fmt"

func main() {
	var celcius int
	var remur int

	fmt.Print("masukan suhu dalam celcius :")
	fmt.Scan(&celcius)

	remur = (celcius * 4) / 5

	fmt.Println("================ HASIL KONVERSI =================")

	fmt.Print("hasil dari konversi remur : ")
	fmt.Println(remur)

	fmt.Println("=================================================")
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguaided/korversi_suhu/output.png)


#### Deskripsi
 program dalam bahasa Go untuk mengonversi suhu dari derajat Celsius menjadi
derajat Reamur dengan menggunakan perkalian (*) dan pembagian (/) untuk menemukan hasil konversi dari celcius ke remur

### 2. hitung harian

```go
package main

import "fmt"

func main() {
	var hari int

	fmt.Print("masukan jumlah hari : ")
	fmt.Scan(&hari)

	// Menggunakan 360 hari untuk 1 tahun (tahun komersial)
	var tahun int = hari / 360
	var sisaHari int = hari % 360

	var bulan int = sisaHari / 30
	sisaHari = sisaHari % 30

	var minggu int = sisaHari / 7
	sisaHari = sisaHari % 7

	fmt.Println("================ HASIL HARI =================")

	fmt.Println("tahun :", tahun)
	fmt.Println("bulan :", bulan)
	fmt.Println("minggu :", minggu)
	fmt.Println("hari :", sisaHari)

	fmt.Println("=============================================")
}
```

##### Output
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguaided/konversi_hari/output.png)

#### Deskripsi
program dalam bahasa Go yang dapat mengonversi jumlah hari ke dalam satuan tahun, bulan, minggu, dan hari menggunakan sisa bagi hasil (%) 

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Kesimpulan diatas saya belajar banyak dari materi Variabel dan Operator yang mulai dari pemrograman konveri suhu, tukar, kasir dan konversi suhu, menghitung hari 

## Referensi
1. King, K. N. (2008). C Programming: A Modern Approach. New York: W. W. Norton & Company. Diakses pada 4 Oktober 2026 melalui ISBN 978-0393979503

2. Donovan, A. A., & Kernighan, B. W. (2015). The Go Programming Language. New York: Addison-Wesley Professional. Diakses pada 4 Oktober 2026 melalui ISBN 978-0134190440

3. Dale, N., & Weems, C. (2012). Programming and Problem Solving with C++. Sudbury: Jones & Bartlett Learning. Diakses pada 4 Oktober 2026 melalui ISBN 978-1449641573

4. Knuth, D. E. (1997). The Art of Computer Programming, Volume 1: Fundamental Algorithms. Reading: Addison-Wesley. Diakses pada 4 Oktober 2026 melalui ISBN 978-0201896831
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
