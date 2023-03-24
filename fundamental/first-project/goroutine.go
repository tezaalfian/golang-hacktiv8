package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	printAcak()
	printRapih()
	wg.Wait()
}

func printAcak() {
	var wg sync.WaitGroup
	wg.Add(8)

	dataStr1 := []interface{}{"Coba1", "Coba2", "Coba3"}
	dataStr2 := []interface{}{"Bisa1", "Bisa2", "Bisa3"}

	// menampilkan secara acak
	for i := 1; i <= 4; i++ {
		i := i
		go func() {
			fmt.Println(dataStr1, i)
			wg.Done()
		}()
		go func() {
			fmt.Println(dataStr2, i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func printRapih() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(8)

	dataStr1 := []interface{}{"Coba1", "Coba2", "Coba3"}
	dataStr2 := []interface{}{"Bisa1", "Bisa2", "Bisa3"}

	//menampilkan secara rapih
	for i := 1; i <= 4; i++ {
		i := i
		go func() {
			mu.Lock()
			fmt.Println(dataStr1, i)
			mu.Unlock()
			wg.Done()
		}()
		go func() {
			mu.Lock()
			fmt.Println(dataStr2, i)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
