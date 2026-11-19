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
		kursUSD = 15800.0
		kursEUR = 17200.0
		kursJPY = 105.0
		kursMYR = 3550.0
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("============================================")
	fmt.Println("    PROGRAM KONVERSI MATA UANG SEDERHANA    ")
	fmt.Println("============================================")

	fmt.Print("Masukkan nominal uang dalam Rupiah (IDR): Rp ")
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("[ERROR] Gagal membaca input.")
		return
	}

	inputStr = strings.TrimSpace(inputStr)
	rupiah, err := strconv.ParseFloat(inputStr, 64)

	// Validasi input: numerik dan harus > 0
	if err != nil {
		fmt.Println("\n[VALIDASI GAGAL] Input harus berupa angka numerik yang valid!")
		return
	}

	if rupiah <= 0 {
		fmt.Println("\n[VALIDASI GAGAL] Nominal uang harus lebih besar dari Rp 0!")
		return
	}

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