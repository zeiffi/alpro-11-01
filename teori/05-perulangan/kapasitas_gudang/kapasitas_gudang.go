package main

import "fmt"

func main() {
    var k, v int
    fmt.Scan(&k)

    totalBarang := 0
    for i := 0; i < 5; i++ {
        fmt.Scan(&v)
        totalBarang += v
        
        if totalBarang >= k {
            fmt.Println(true)
        } else {
            fmt.Println(false)
        }
    }
}