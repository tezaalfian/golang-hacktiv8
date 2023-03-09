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

	str := "CIAIШAIPIBIО"
	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			fmt.Printf("character U+%04X '%c' starts at byte position %d\n", str[i], str[i], i)
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
