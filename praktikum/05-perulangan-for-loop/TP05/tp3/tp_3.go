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