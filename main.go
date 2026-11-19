package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Kurs tetap mata uang terhadap Rupiah (IDR)
	const (
		kursUSD = 15800.0 // 1 USD = Rp 15.800
		kursEUR = 17200.0 // 1 EUR = Rp 17.200
		kursJPY = 105.0   // 1 JPY = Rp 105
		kursMYR = 3550.0  // 1 MYR = Rp 3.550

		persenAdmin = 0.01 // Fitur Pengembangan: Biaya admin penukaran uang 1%
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("====================================================")
	fmt.Println("         PROGRAM KONVERSI MATA UANG SEDERHANA       ")
	fmt.Println("             DENGAN FITUR BIAYA ADMIN 1%            ")
	fmt.Println("====================================================")

	fmt.Print("Masukkan nominal uang dalam Rupiah (IDR): Rp ")
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("[ERROR] Gagal membaca input pengguna.")
		return
	}

	inputStr = strings.TrimSpace(inputStr)
	rupiahAwal, err := strconv.ParseFloat(inputStr, 64)

	// Validasi input
	if err != nil {
		fmt.Println("\n[VALIDASI GAGAL] Input harus berupa angka numerik yang valid!")
		return
	}

	if rupiahAwal <= 0 {
		fmt.Println("\n[VALIDASI GAGAL] Nominal uang harus lebih besar dari Rp 0!")
		return
	}

	// Fitur Pengembangan: Pemotongan Biaya Admin 1%
	biayaAdmin := rupiahAwal * persenAdmin
	rupiahBersih := rupiahAwal - biayaAdmin

	// Perhitungan konversi dari nominal bersih
	usd := rupiahBersih / kursUSD
	eur := rupiahBersih / kursEUR
	jpy := rupiahBersih / kursJPY
	myr := rupiahBersih / kursMYR

	// Tampilkan rincian transaksi & hasil konversi
	fmt.Println("\n---------------- RINCIAN TRANSAKSI -----------------")
	fmt.Printf("Nominal Awal (Kotor) : Rp %15.2f\n", rupiahAwal)
	fmt.Printf("Biaya Admin (1%%)     : Rp %15.2f\n", biayaAdmin)
	fmt.Printf("Nominal Bersih       : Rp %15.2f\n", rupiahBersih)

	fmt.Println("\n------------------ HASIL KONVERSI ------------------")
	fmt.Printf("1. Dollar AS  (USD)  : $  %15.2f  (Kurs: Rp %.0f)\n", usd, kursUSD)
	fmt.Printf("2. Euro       (EUR)  : €  %15.2f  (Kurs: Rp %.0f)\n", eur, kursEUR)
	fmt.Printf("3. Yen Jepang (JPY)  : ¥  %15.2f  (Kurs: Rp %.0f)\n", jpy, kursJPY)
	fmt.Printf("4. Ringgit    (MYR)  : RM %15.2f  (Kurs: Rp %.0f)\n", myr, kursMYR)
	fmt.Println("====================================================")
}