package main

import "fmt"

func main() {
	var a,b int

	//membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//menukar a dan b
	a,b=b,a

	//output
	fmt.Println(a)
	fmt.Println(b)
}