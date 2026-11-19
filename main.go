package main

import (
	"fmt"
)

func main() {
	// Kurs tetap mata uang terhadap Rupiah (IDR)
	const (
		kursUSD = 15800.0 // 1 USD = Rp 15.800
		kursEUR = 17200.0 // 1 EUR = Rp 17.200
		kursJPY = 105.0   // 1 JPY = Rp 105
		kursMYR = 3550.0  // 1 MYR = Rp 3.550
	)

	var rupiah float64

	fmt.Println("============================================")
	fmt.Println("    PROGRAM KONVERSI MATA UANG SEDERHANA    ")
	fmt.Println("============================================")

	fmt.Print("Masukkan nominal uang dalam Rupiah (IDR): Rp ")
	fmt.Scanln(&rupiah)

	// Perhitungan konversi
	usd := rupiah / kursUSD
	eur := rupiah / kursEUR
	jpy := rupiah / kursJPY
	myr := rupiah / kursMYR

	fmt.Println("\n---------------- Hasil Konversi ------------")
	fmt.Printf("Nominal IDR : Rp %.2f\n", rupiah)
	fmt.Printf("1. Dollar AS  (USD) : $ %.2f\n", usd)
	fmt.Printf("2. Euro       (EUR) : € %.2f\n", eur)
	fmt.Printf("3. Yen Jepang (JPY) : ¥ %.2f\n", jpy)
	fmt.Printf("4. Ringgit    (MYR) : RM %.2f\n", myr)
	fmt.Println("============================================")
}