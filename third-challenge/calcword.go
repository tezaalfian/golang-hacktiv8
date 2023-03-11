package main

import (
	"fmt"
)

func main() {
	kata := "selamat malam"

	// Inisialisasi map untuk menyimpan jumlah kemunculan setiap huruf
	kataCount := make(map[string]int)

	// Hitung jumlah kemunculan setiap huruf
	for _, c := range kata {
		if string(c) == " " {
			continue
		}
		kataCount[string(c)]++
	}

	fmt.Print(kataCount)
}
