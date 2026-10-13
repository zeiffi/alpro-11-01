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