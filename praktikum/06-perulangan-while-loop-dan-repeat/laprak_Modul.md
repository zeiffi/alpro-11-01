# <h1 align="center">Laporan Praktikum Modul 06 - Perulangan While</h1>
<p align="center">MUHAMMAD SHEIF FAHRI - 109092600016</p>

## Dasar Teori

### A. Konsep Dasar Perulangan Bersyarat (While)
Perulangan bersyarat atau yang secara umum dikenal sebagai while-loop adalah struktur kendali alur (control flow) dalam pemrograman yang mengeksekusi blok kode secara berulang selama kondisi ekspresi logika bernilai benar (true). Pengecekan kondisi dilakukan di awal sebelum blok pernyataan dijalankan. Jika kondisi bernilai salah (false) sejak awal, maka blok perulangan tidak akan pernah dieksekusi sama sekali. Struktur ini sangat ideal digunakan untuk kasus yang jumlah perulangannya tidak pasti di awal (indefinite iteration), seperti proses validasi input login ATM, pembacaan file baris per baris, atau pencarian nilai matematika tertentu hingga batas ambang tercapai.
### B. Implementasi Perulangan pada Bahasa Go (Golang)

#### 1. Studi Paradigma Bahasa Pemrograman:
Berdasarkan analisis performa dan sintaksis bahasa Go yang dipublikasikan dalam literatur rekayasa perangkat lunak modern, desain minimalis dengan menyatukan seluruh pola perulangan ke dalam satu kata kunci for bertujuan untuk mengurangi kompleksitas leksikal (lexical complexity) dan mempermudah pemeliharaan kode (maintainability) pada sistem berskala besar.

#### 2. Keamanan dan Validasi Akses (Studi Kasus ATM):
Dalam standar perancangan sistem autentikasi tertutup, penerapan batas percobaan (maximum retry limit) menggunakan perulangan bersyarat merupakan implementasi standar keamanan lapis pertama untuk mencegah serangan brute-force pada terminal perbankan mini atau simulasi ATM, sebagaimana distandarisasi dalam literatur pengelolaan integritas sistem informasi.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. terkecil.go

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n*n <= batas {
        n++
    }

    fmt.Println(n)
}
```

#### Deskripsi

Program Go ini membaca sebuah bilangan bulat batas dari input, lalu mencari bilangan bulat terkecil n yang pangkat tiganya (n*n*n) sudah melebihi batas, dengan cara menaikkan n mulai dari 1 selama n³ masih kurang dari atau sama dengan batas, dan hasilnya dicetak di akhir. Misalnya, jika inputnya 20, outputnya 3 karena 1³ = 1 dan 2³ = 8 masih ≤ 20, sedangkan 3³ = 27 sudah melebihi 20.

### 2. terbesar.go

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n < batas {
        n++
    }

    fmt.Println(n - 1)
}
```
#### Deskripsi

Program Go ini membaca sebuah bilangan bulat batas dari input, lalu menaikkan n mulai dari 1 selama n*n masih kurang dari batas, dan mencetak n - 1 di akhir. Hasilnya adalah bilangan bulat terbesar yang kuadratnya masih kurang dari batas (strictly lebih kecil). Misalnya, jika inputnya 20, outputnya 4 karena 4² = 16 < 20 sedangkan 5² = 25 sudah tidak kurang dari 20. Perlu dicatat, jika batas sendiri berupa kuadrat sempurna (misalnya 16), outputnya 3, bukan 4, karena 4² = 16 tidak kurang dari 16.


### 3. vokal.go

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    kalimat, _ := reader.ReadString('\n')

    i := 0
    jumlahVokal := 0

    for i < len(kalimat) {
        ch := strings.ToLower(string(kalimat[i]))
        if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
            jumlahVokal++
        }
        i++
    }

    fmt.Println(jumlahVokal)
}
```
#### Deskripsi
Program Go ini membaca satu baris kalimat dari input pengguna, lalu menghitung jumlah huruf vokal (a, i, u, e, o) di dalamnya dengan mengecek setiap karakter satu per satu menggunakan perulangan for. Setiap karakter diubah ke huruf kecil terlebih dahulu agar huruf kapital juga ikut terhitung, dan hasil totalnya dicetak di akhir program.


<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. ATM

```go
package main

import (
	"fmt"
)

func main() {
	var password string
	var kesempatan int = 3
	var loginBerhasil bool = false

	for kesempatan > 0 && !loginBerhasil {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password == "1234" {
			fmt.Println("Login berhasil!")
			loginBerhasil = true
		} else {
			kesempatan--
			if kesempatan > 0 {
				fmt.Println("Password salah. Silakan coba lagi.")
			} else {
				fmt.Println("Password salah.")
			}
		}
	}

	if !loginBerhasil {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	var saldo int = 10000000
	var pilihan int

	for pilihan != 3 {
		fmt.Println("\nMenu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilihan == 2 {
			var tarik int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&tarik)

			if tarik <= saldo {
				saldo -= tarik
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", tarik)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/unguided/output.png)


#### Deskripsi
 wow Pemrograman adalah proses menulis, menguji, dan memelihara serangkaian instruksi terstruktur menggunakan bahasa komputer tertentu agar dapat dijalankan oleh perangkat untuk menyelesaikan tugas atau memecahkan masalah secara otomatis.


## Kesimpulan
Berdasarkan praktikum yang telah dilakukan, dapat disimpulkan bahwa bahasa Go meskipun tidak memiliki kata kunci while secara eksplisit, tetap dapat mengimplementasikan perulangan bersyarat menggunakan sintaks for dengan menyertakan kondisi logika tunggal. Melalui penerapan struktur ini, program simulasi ATM dapat berjalan secara interaktif—mengontrol pembatasan jumlah percobaan kesalahan autentikasi (retry limit) serta mempertahankan siklus menu transaksi hingga pengguna memilih untuk keluar dari sistem.

## Referensi
1. Donovan, A. A. A., & Kernighan, B. W. (2015). The Go Programming Languages. 
2. Sumarni, L., & Wijaya, A. (2022). Analisis Kompleksitas Sintaksis dan Performa Struktur Kontrol Alur pada Bahasa Pemrograman Modern.
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
