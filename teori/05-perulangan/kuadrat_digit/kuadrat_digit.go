package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)

    var kuadratStr string
    var total int

    for x > 0 {
        digit := x % 10
        k := digit * digit
        total += k
        
        if kuadratStr == "" {
            kuadratStr = fmt.Sprintf("%d", k)
        } else {
            kuadratStr = fmt.Sprintf("%d %s", k, kuadratStr)
        }
        
        x = x / 10
    }

    fmt.Println(kuadratStr)
    fmt.Println(total)
}