# <h1 align="center">Tugas Pendahuluan Modul 05 - Perulangan For</h1>
<p align="center">MUHAMMAD SEHIF FAHRI - 109092600016</p>

### 1. soal
Apa fungsi utama dari struktur perulangan for dalam program Go? Sebutkan dan
jelaskan tiga komponen pada perulangan for (inisialisasi, kondisi, dan post statement)!

jawab :
    Fungsi Utama: Digunakan untuk mengeksekusi blok kode secara berulang kali secara efisien berdasarkan kondisi atau aturan tertentu.

Tiga Komponen:

1. Inisialisasi (Initialization): Pernyataan awal untuk mendeklarasikan dan memberi nilai awal pada variabel penghitung (counter) sebelum perulangan dimulai (contoh: i := 1). Komponen ini hanya dieksekusi sekali di awal.

2. Kondisi (Condition): Ekspresi boolean yang dievaluasi sebelum setiap iterasi. Jika bernilai true, blok kode di dalam perulangan akan dijalankan. Jika false, perulangan berhenti.

3. Post Statement: Pernyataan yang dieksekusi di setiap akhir iterasi (biasanya digunakan untuk menaikkan atau menurunkan nilai variabel penghitung, contoh: i++).


### 2. soal
Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara
for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh
kode masing-masing.

Penjelasan Perbedaan:

- for dengan tiga komponen (lengkap) biasanya digunakan ketika jumlah iterasi atau perulangan sudah diketahui dengan jelas sejak awal (misalnya dari 1 sampai $n$)

- for yang hanya menggunakan kondisi saja (mirip konsep while di bahasa lain) digunakan ketika perulangan bergantung pada kondisi tertentu di mana kita tidak tahu pasti berapa kali perulangan akan berjalan, melainkan berhenti saat kondisi tersebut bernilai false.

Contoh Kode:

- For Tiga Komponen:
```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```
- For Kondisi Saja (Mirip While):
```go
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
```


### 3. ganjil.go

```go
package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan nilai n: ")
    fmt.Scan(&n)

    var jumlah int = 0
    fmt.Printf("Bilangan ganjil dari 1 hingga %d: ", n)

    for i := 1; i <= n; i++ {
        if i%2 != 0 {
            fmt.Printf("%d ", i)
            jumlah += i
        }
    }
    fmt.Printf("\nJumlah = %d\n", jumlah)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/TP05/tp3/output.png)


#### Deskripsi
Program ini berfungsi untuk menerima masukan sebuah bilangan bulat positif $n$, lalu menggunakan perulangan for untuk menyaring dan menampilkan seluruh bilangan ganjil dari angka 1 hingga $n$. Selain menampilkan deret bilangan ganjilnya, program ini juga menjumlahkan seluruh bilangan tersebut dan mencetak total akhirnya.

### 4. perkalian.go

```go
package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("Masukkan nilai a dan b: ")
    fmt.Scan(&a, &b)

    fmt.Printf("Tabel perkalian %d, dari %d x 1 sampai %d x %d:\n", a, a, a, b)
    for i := 1; i <= b; i++ {
        hasil := a * i
        fmt.Printf("%d x %d = %d\n", a, i, hasil)
    }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/TP05/tp4/output.png)


#### Deskripsi
Program ini dirancang untuk menerima dua buah masukan bilangan bulat positif, yaitu $a$ (angka yang akan dikalikan) dan $b$ (batas akhir perkalian). Menggunakan perulangan for, program ini mencetak tabel perkalian mulai dari $a \times 1$ secara berurutan ke bawah hingga mencapai batas $a \times b$ dengan format yang rapi.


## Kesimpulan
Struktur perulangan for dalam bahasa Go sangat efisien dan fleksibel untuk mengeksekusi blok kode secara berulang—baik menggunakan tiga komponen lengkap maupun berbasis kondisi saja—sehingga sangat ideal dipadukan dengan percabangan dan pengolahan parameter ganda untuk menyaring data, mengakumulasi nilai, serta menampilkan tabel terstruktur seperti pada program ganjil.go dan perkalian.go.