package main

import "fmt"

func awal() {
	for i := 10; i > 0; i-- {
		defer func() {
			fmt.Println("Awal mulai")
		}()
	}

}
func kedua() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println("Angka", i)
		}()
	}
	fmt.Println("Kedua selesai")
}

func ketiga() {
	for i := 10; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	fmt.Println("Ketiga selesai")
}

func main() {
	awal()
	kedua()

	ketiga()
}
