package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)

    cantik := true
    prevDigit := 10

    for x > 0 {
        currentDigit := x % 10
        if currentDigit > prevDigit {
            cantik = false
            break
        }
        prevDigit = currentDigit
        x = x / 10
    }

    fmt.Println(cantik)
}