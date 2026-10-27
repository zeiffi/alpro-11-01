# <h1 align="center">Tugas Pendahuluan Modul 07 - PERCABANGAN</h1>
<p align="center">MUHAMMAD SHEIF FAHRI - 109092600016</p>

### 1. perbedaan switc pada bahasa golang dan pada bahasa java

1. Perbedaan Struktur Percabangan & Switch
if: Digunakan untuk mengeksekusi satu blok kode jika suatu kondisi bernilai true.

if-else: Menyeleksi dua kondisi; jika kondisi if bernilai false, maka blok else yang akan dijalankan.

else if: Digunakan untuk menguji lebih dari dua kondisi secara berurutan.

switch: Digunakan untuk memilih salah satu dari banyak blok kode yang akan dieksekusi berdasarkan pencocokan nilai ekspresi.

2.  Perbedaan Switch pada Go vs Java:

Implisit break: Di Go, setiap case secara otomatis berhenti (memiliki break implisit) sehingga tidak perlu menuliskan kata kunci break. Di Java, jika tidak ada break, eksekusi akan lanjut ke case bawahnya (fallthrough).

Tanpa Ekspresi (Tanpa Variabel): Di Go, switch bisa ditulis tanpa menyertakan variabel di awal (misal: switch { case x > 0: ... }), sehingga case dapat berisi kondisi boolean kompleks seperti if-else. Di Java, switch wajib menerima nilai/variabel tertentu.


### 2. bilangan.go

```go
package main

import "fmt"

func main() {
    var n int

	fmt.Print("masukan bilangan : ")
    fmt.Scan(&n)

    if n == 0 {
        fmt.Println("Nol")
    } else {
        var statusGenap, statusPositif string

        if n%2 == 0 {
            statusGenap = "Genap"
        } else {
            statusGenap = "Ganjil"
        }

        if n > 0 {
            statusPositif = "Positif"
        } else {
            statusPositif = "Negatif"
        }
		
        fmt.Printf("%s, %s\n", statusGenap, statusPositif)
    }
}	
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/07-percabangan/TP/TP-1/output.png)


#### Deskripsi
memasukan bilangan dengan output menentukan bilangan genap atau ganjil dan positif atau negatif

### 3. hari.go

```go
package main

import "fmt"

func main() {
    var hari int

	fmt.Print("masukan hari : ")
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
        fmt.Println("Sabtu (akhir pekan)")
    case 7:
        fmt.Println("Minggu (akhir pekan)")
    default:
        fmt.Println("Nomor hari tidak valid")
    }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/07-percabangan/TP/TP-2/output.png)


#### Deskripsi
Program menerima input angka 1–7 dan mencetak nama hari menggunakan switc

### 2. bilangan.go

```go
package main

import "fmt"

func main() {
    var a, b, c int

	fmt.Print("masukan 3 bilangan : ")
    fmt.Scan(&a, &b, &c)

    if a >= b && a >= c {
        fmt.Println(a)
    } else if b >= a && b >= c {
        fmt.Println(b)
    } else {
        fmt.Println(c)
    }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/07-percabangan/TP/TP-3/output.png)


#### Deskripsi
Program menerima 3 bilangan bulat dan mencari mana yang paling besar tanpa perulangan atau fungsi max


## Kesimpulan
Secara keseluruhan, keempat tugas ini dirancang untuk menguasai logika percabangan di Go; dimulai dari pemahaman konseptual struktur kontrol (if, if-else, else if, switch) pada Soal 1, dilanjutkan dengan penerapan logika kondisi majemuk untuk menentukan sifat bilangan pada bilangan.go, penggunaan switch-case untuk pemetaan kategori hari pada hari.go, hingga perbandingan multi-kondisi berjenjang guna mencari nilai terbesar pada terbesar.go.