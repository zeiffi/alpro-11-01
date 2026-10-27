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