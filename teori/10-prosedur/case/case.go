package main

import "fmt"

func add(a,b int32) int32 { //fungsi punya keyword tipe data kembali dan return 
	return a + b
}

func eating() {
	fmt.Println("saya lapar")
	fmt.Println("saya mengambil piring")
	fmt.Println("saya mengambil nais dan lauk")
	fmt.Println("saya makan nasi dan lauk")
}

func main() {
	fmt.Println(add(2, 3))
	eating()
}