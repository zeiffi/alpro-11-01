package main

import (
	"fmt"
)

func main() {
	var password string
	var kesempatan int = 3
	var loginBerhasil bool = false

	for kesempatan > 0 && !loginBerhasil {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password == "1234" {
			fmt.Println("Login berhasil!")
			loginBerhasil = true
		} else {
			kesempatan--
			if kesempatan > 0 {
				fmt.Println("Password salah. Silakan coba lagi.")
			} else {
				fmt.Println("Password salah.")
			}
		}
	}

	if !loginBerhasil {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	var saldo int = 10000000
	var pilihan int

	for pilihan != 3 {
		fmt.Println("\nMenu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilihan == 2 {
			var tarik int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&tarik)

			if tarik <= saldo {
				saldo -= tarik
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", tarik)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}