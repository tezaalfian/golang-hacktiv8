package main

import "fmt"

func main() {
	// menampilkan nilai i : 21
	i := 21
	fmt.Printf("%v \n", i)
	// menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)
	// menampilkan tanda %
	fmt.Printf("%%\n")
	// menampilkan nilai boolean j : true
	j := true
	fmt.Printf("%v \n", j)
	// menampilkan unicode russia : Я (ya)
	fmt.Printf("%c\n", '\u042F')
	// menampilkan nilai base 10 : 21
	fmt.Printf("%d \n", 21)
	// menampilkan nilai base 8 :25
	fmt.Printf("%o \n", 21)
	// menampilkan nilai base 16 : f
	fmt.Printf("%x \n", 15)
	// menampilkan nilai base 16 : F 13
	fmt.Printf("%X \n", 15)
	// menampilkan unicode karakter Я : U+042F
	fmt.Printf("%U \n", '\u042F')
	// menampilkan float : 123.456000
	var k float64 = 123.456
	fmt.Printf("%f \n", k)
	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%e \n", k)
}
