# <h1 align="center">Laporan Praktikum Modul 04 - Runtunan/Sekuensi</h1>
<p align="center">MUHAMMAD SHEIF FAHRI - 109092600016</p>

## Dasar Teori

### A. Konsep Dasar Struktur Kontrol Percabangan (Branching) dalam Bahasa Go
Struktur kontrol percabangan adalah mekanisme dalam pemrograman yang memungkinkan alur eksekusi program berjalan secara tidak linier, melainkan bercabang berdasarkan evaluasi suatu kondisi bernilai boolean (true atau false). Menurut Munir (2020), struktur pemilihan atau percabangan memegang peranan krusial dalam algoritma untuk mengambil keputusan berdasarkan kriteria tertentu sehingga program dapat memberikan hasil yang dinamis sesuai dengan data masukan. Dalam bahasa pemrograman Go (Golang), percabangan diimplementasikan melalui konstruksi logika seperti if, else, dan switch yang dirancang untuk mengevaluasi ekspresi kondisional secara efisien dan terstruktur.

### B. Implementasi Percabangan if-else dan switch pada Bahasa Pemrograman Go

#### 1. Pernyataan Percabangan if, else if, dan else
Pernyataan if digunakan untuk mengevaluasi suatu kondisi tunggal maupun jamak. Jika kondisi bernilai benar (true), maka blok kode di dalam if akan dieksekusi; jika salah, alur program dapat melanjutkan pengecekan ke kondisi berikutnya melalui else if atau mengeksekusi blok alternatif pada else. Dalam bahasa Go, tanda kurung () di sekitar ekspresi kondisi bersifat opsional, namun penggunaan kurung kurawal {} wajib disertakan untuk mendefinisikan blok kode. Percabangan berjenjang (else if) atau bersarang sangat efektif digunakan dalam penyelesaian kasus yang memerlukan pengujian rentang nilai bertingkat, seperti sistem penentuan indeks nilai akademik maupun perhitungan bracket pajak progresif.

#### 2. Pernyataan Percabangan switch-case
Pernyataan switch di dalam bahasa Go merupakan bentuk alternatif dari percabangan if-else bertingkat yang membuat struktur kode menjadi jauh lebih bersih, rapi, dan mudah dibaca saat menghadapi banyak pilihan kondisi. Berbeda dengan beberapa bahasa pemrograman konvensional, pernyataan switch di Go secara otomatis menghentikan eksekusi (implicit break) setelah sebuah case yang cocok terpenuhi, sehingga tidak memerlukan kata kunci break secara manual. Selain itu, Go mendukung fitur tagless switch (switch tanpa ekspresi variabel di awal), di mana setiap case dapat langsung berisi kondisi boolean independen, memberikan fleksibilitas tinggi untuk pengelompokan klasifikasi data yang kompleks.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. grade.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scanln(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "F"
	}
	fmt.Printf("%s mendapatkan nama %s\n", nama, grade)
}
```
#### Deskripsi
pemrograman ini menentukan nilai grade berdasarkan nilai angka yang dimasukkan oleh pengguna. Program menerima input berupa nama dan nilai, kemudian menggunakan percabangan if-else untuk menentukan grade: A (90–100), B (80–89), C (70–79), D (60–69), dan F (di bawah 60). Setelah itu, program menampilkan nama dan grade yang diperoleh.

### 2. penilaian.go

```go
package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
}
```
#### Deskripsi
Program ini menggunakan bahasa Go untuk membuat sistem penilaian siswa berbasis menu. Pengguna dapat memilih untuk memasukkan nama dan nilai siswa, kemudian program menentukan grade A, B, C, D, atau F berdasarkan nilai yang diperoleh. Jika pengguna memilih 0, program akan menampilkan pesan keluar, sedangkan pilihan lain yang tidak valid akan menampilkan pesan kesalahan.

### 3. klasifikasi.go

```go
package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}
```
#### Deskripsi
pemrograman ini untuk mengklasifikasikan seseorang berdasarkan usia dan gaji yang dimasukkan oleh pengguna. Dengan menggunakan percabangan switch, program menentukan keterangan yang sesuai, seperti masih sekolah, muda sukses, pekerja mapan, atau profesional berpengalaman, berdasarkan rentang usia dan jumlah gaji. Hasil klasifikasi kemudian ditampilkan ke layar.


<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. gade

```go
package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    var nama string
    var nilai float64
    var grade string

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Masukkan nama: ")
    scanner.Scan()
    nama = scanner.Text()

    fmt.Print("Masukkan nilai: ")
    fmt.Scan(&nilai)

    switch {
    case nilai >= 90 && nilai <= 100:
        grade = "A"
    case nilai >= 80 && nilai < 90:
        grade = "B"
    case nilai >= 70 && nilai < 80:
        grade = "C"
    case nilai >= 60 && nilai < 70:
        grade = "D"
    default:
        grade = "f"
    }

    fmt.Printf("%s, mendapatkan nilai %s\n", nama, grade)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/grade/output.png)


#### Deskripsi
Program ini menggunakan bahasa Go untuk menentukan grade nilai siswa berdasarkan nama dan nilai yang dimasukkan. Dengan menggunakan percabangan switch, program mengelompokkan nilai ke dalam grade A, B, C, D, atau F, kemudian menampilkan nama siswa beserta grade yang diperoleh.

### 2. pajak

```go
package main

import "fmt"

func main() {

	var penghasilan, pajak float64
	fmt.Print("Masukkan penghasilan: ")
	fmt.Scan(&penghasilan)

if penghasilan <= 50 {
    pajak = penghasilan * 0.05
} else if penghasilan <= 100 {
    pajak = (50 * 0.05) + ((penghasilan - 50) * 0.10)
} else if penghasilan <= 200 {
    pajak = (50 * 0.05) + (50 * 0.10) + ((penghasilan - 100) * 0.15)
} else {
    pajak = (50 * 0.05) + (50 * 0.10) + (100 * 0.15) + ((penghasilan - 200) * 0.20)
}

fmt.Println("Pajak:", pajak)
}
```

##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/pajak/output.png)

#### Deskripsi
pemograman ini untuk menghitung pajak berdasarkan jumlah penghasilan yang dimasukkan oleh pengguna. Dengan menggunakan percabangan if-else, program menerapkan tarif pajak bertingkat, yaitu 5%, 10%, 15%, dan 20% sesuai rentang penghasilan. Setelah perhitungan selesai, program menampilkan total pajak yang harus dibayarkan.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Secara keseluruhan, meskipun modul ini berjudul sekuensial, inti dari praktikum bahasa Go ini adalah melatih kemampuan logika percabangan menggunakan if-else dan switch—seperti pada program penilaian, klasifikasi usia, hingga perhitungan pajak—agar kita dapat membangun program yang dinamis, terstruktur, dan mampu mengambil keputusan secara otomatis berdasarkan kondisi data.

## Referensi
1. Donovan, A. A. A., & Kernighan, B. W. (2015) https://www.gopl.io/

2. Munir, R. (2020)Bandung: Informatika. Diakses pada 11 Oktober 2026 melalui direktori literatur akademik Informatika
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
