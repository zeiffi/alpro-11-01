# <h1 align="center">Tugas Pendahuluan Modul 04 - RUNTUNAN/SEKUENSI</h1>
<p align="center">MUHAMMAD SEHIF FAHRI - 109092600016</p>

### 1. Evaluasi Ekspresi Kontrol dalam Go

```go
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	fmt.Println(intNum > 5 && intOther > 5 && sngNum > 0)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP_04/tp_1/output.png)


#### Deskripsi
Program di atas adalah program sederhana yang ditulis menggunakan bahasa Go (Golang). Program ini mendemonstrasikan penggunaan variabel dengan berbagai tipe data dan evaluasi ekspresi logika menggunakan operator && (AND).
### 2.  Tracing: Evaluasi Pernyataan Kondisi

```go
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}

	fmt.Println("Nilai akhir result:", result)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP_04/tp_2/output.png)


#### Deskripsi
Program di atas adalah program dalam bahasa Go (Golang) yang menggunakan beberapa variabel bertipe integer (x, y, z, dan result) untuk melakukan serangkaian operasi aritmatika yang dikontrol oleh struktur percabangan kondisi (if-else).

### 3.  Menentukan Jumlah Hari dalam Sebulan Berdasarkan Tahun dan Bulan

```go
package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun, &bulan)

	// Cek tahun kabisat
	kabisat := (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0

	var hari int

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		hari = 31
	case "Apr", "Jun", "Sep", "Nov":
		hari = 30
	case "Feb":
		if kabisat {
			hari = 29
		} else {
			hari = 28
		}
	default:
		// Nama bulan tidak valid: tidak ada keluaran ("-")
		fmt.Println("-")
		return
	}

	fmt.Println(hari)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP_04/tp_4/output.png)


#### Deskripsi
Program Go tersebut berfungsi untuk menentukan jumlah hari dalam suatu bulan berdasarkan input tahun dan singkatan nama bulan dari pengguna, di mana program akan mengecek tahun kabisat terlebih dahulu untuk menghitung jumlah hari bulan Februari secara akurat (28 atau 29 hari), mengelompokkan bulan-bulan berhari 30 dan 31, mencetak tanda strip (-) jika nama bulan tidak valid, lalu menampilkan hasil akhirnya ke layar.

### 4.  Switch Case

```go
package main

import "fmt"

func main() {
	var hari int
	fmt.Print("Masukkan angka hari (1-7): ")
	fmt.Scan(&hari)

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Hari tidak valid")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP_04/tp_34/output.png)


#### Deskripsi
Program Go tersebut meminta pengguna memasukkan angka dari 1 hingga 7, lalu menggunakan struktur switch-case untuk mencetak nama hari yang sesuai (Senin sampai Minggu) atau menampilkan pesan "Hari tidak valid" jika angka yang dimasukkan berada di luar rentang tersebut.

## Kesimpulan
Jadi, intinya, semua contoh kode tadi tuh semacam latihan dasar yang pas banget buat ngenalin kita ke cara kerja pemrograman Go—mulai dari bikin variabel, mainan tipe data, nyoba operator logika sama matematika, sampai pakai percabangan if-else dan switch-case buat ngolah data inputan dari user dan nampilin hasilnya secara akurat.