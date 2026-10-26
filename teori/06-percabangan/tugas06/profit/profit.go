package main

import "fmt"

func main() {
	var bulanIni, bulanLalu float64

	fmt.Print("Masukkan profit bulan lalu: ")
	fmt.Scanln(&bulanLalu)
	fmt.Print("Masukkan profit bulan ini: ")
	fmt.Scanln(&bulanIni)

	if bulanIni > bulanLalu {
		sisa := bulanIni - bulanLalu
		fmt.Printf("Peningkatan sebesar %.2f\n", sisa)
	} else if bulanIni < bulanLalu {
		sisa := bulanIni - bulanLalu
		fmt.Printf("Penurunan sebesar %.2f\n", sisa)
	} else {
		fmt.Println("Tetap")
	}
}

// program ProfitBulanan
// kamus
//     bulanIni, bulanLalu, sisa : real
// 
// algoritma
//     output("Masukkan profit bulan lalu: ")
//     input(bulanLalu)
//     output("Masukkan profit bulan ini: ")
//     input(bulanIni)
// 
//     if bulanIni > bulanLalu then
//         sisa ← bulanIni - bulanLalu
//         output("Peningkatan sebesar ", sisa)
//     else if bulanIni < bulanLalu then
//         sisa ← bulanIni - bulanLalu
//         output("Penurunan sebesar ", sisa)
//     else
//         output("Tetap")
//     endif
// endprogram