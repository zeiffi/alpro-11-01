package main
 
import "fmt"
 
func main() {
	fmt.Println("Masukkan tipe kendaraan (M/K/T), tekan enter untuk selesai:")
	
	var c string
	var m, k, t int
 
	for {
		n, _ := fmt.Scan(&c)
		if n != 1 {
			break
		}
		if c == "M" {
			m++
		} else if c == "K" {
			k++
		} else if c == "T" {
			t++
		} else {
			break 
		}
	}
 
	fmt.Println("Jumlah kendaraan yang dimasukkan:")
	fmt.Println("Tipe M =", m)
	fmt.Println("Tipe K =", k)
	fmt.Println("Tipe T =", t)
}
 