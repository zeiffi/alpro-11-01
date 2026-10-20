package main

import "fmt"

func main() {
    var poin, total int
    total = 0

	fmt.Println("masukan poin dan diakiri 0 ")
    fmt.Scan(&poin)
    for poin != 0 {
        total += poin
        fmt.Scan(&poin)
    }
	fmt.Println("Total poin yang didapatkan : ")
    fmt.Println(total)
}