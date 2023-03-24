package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	j := 0
	for j < 5 {
		fmt.Println("Nilai j = ", j)
		j++
	}

	str := "C A Ш A P B О"
	runeStr := []rune(str)
	for i := 0; i < len(runeStr); i++ {
		if i%2 == 0 {
			fmt.Printf("character %U '%c' starts at byte position %d\n", runeStr[i], runeStr[i], i)
		}
	}

	j = 6
	for {
		fmt.Printf("Nilai j = %d \n", j)
		j++
		if j == 11 {
			break
		}
	}
}
