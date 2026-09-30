# <h1 align="center">Tugas Pendahuluan Modul 03 - Tipe Data Dan Intruksi Dasar</h1>
<p align="center">MUHAMMAD SHEIF FAHRI - 19092600016</p>

### 1.  Sisa Kue

```go
package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("masukan jumlah kue	:")
	fmt.Scanln(&x)

	fmt.Print("masukan jumlah kue	:")
	fmt.Scanln(&y)

	fmt.Println("hasil kue sisa bagi	:", x % y)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](TP_1/Output.png)


#### Deskripsi
menghitung hasil bagi sisa kue yang  berupa dua buah bilangan bulat y dan x yang menyatakan jumlah kue dan 
jumlah anggota keluarga. berupa jumlah kue yang tersisa setelah kue dibagikan sama rata kepada 
setiap anggota keluarga. 

### 2. bool

```go
package main

import "fmt"

func main() {
	var x, y bool

	fmt.Print("masukan nilai ke1	: ")
	fmt.Scan(&x)
	fmt.Print("masukan nilai ke2	: ")
	fmt.Scan(&y)

	fmt.Print("hasil ke1	: ")
	fmt.Println(x)
	fmt.Print("hasil ke	: ")
	fmt.Println(y)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/TP_2/Output.png)


#### Deskripsi
menentukan nilai true atau false menggunakan boolean

### 3. konversi

```go
package main

import "fmt"

func main (){
	var x float64
	var y float64 = 1.6

	fmt.Print("masukan mil :")
	fmt.Scanln(&x)

	y = x * y
	
	fmt.Print("hasil km :")
	fmt.Printf( "%.1f", y)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/TP_3/Output.png)


#### Deskripsi 
mengonversikan niai dari mill ke kilometer 1 mil = 1.6 kilometer tampilkan hasil konversi dengan 1 angka di belakang koma (gunakan 
fmt.Printf dengan format %.1f).

## Kesimpulan
saya dapat belajar dari tugas TP di atas dengan mempelajari sisa bagi hasil, boolean, fmt.Printf dengan format %.1f