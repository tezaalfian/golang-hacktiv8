package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	No        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var temanList = []Teman{
	{1, "John Doe", "Jakarta", "Software Engineer", "Ingin belajar bahasa pemrograman baru"},
	{2, "Jane Doe", "Surabaya", "Data Analyst", "Ingin mengembangkan keterampilan teknis"},
	{3, "Bob Smith", "Bali", "UI/UX Designer", "Ingin belajar teknologi baru"},
	{4, "Alice Brown", "Bandung", "Project Manager", "Ingin meningkatkan produktivitas kerja"},
	{5, "David Lee", "Yogyakarta", "Business Analyst", "Ingin memperluas pengetahuan bisnis"},
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run biodata.go <no absen>")
		return
	}

	no, err := strconv.Atoi(args[1])
	if err != nil || no < 1 || no > len(temanList) {
		fmt.Println("Invalid argument. Please enter a valid number.")
		return
	}

	teman := temanList[no-1]
	fmt.Printf("No Absen: %d\n", teman.No)
	fmt.Printf("Nama: %s\n", teman.Nama)
	fmt.Printf("Alamat: %s\n", teman.Alamat)
	fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
	fmt.Printf("Alasan memilih kelas Golang: %s\n", teman.Alasan)
}
