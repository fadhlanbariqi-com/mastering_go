// Garbage Collector

package main

import (
	"fmt"
	"runtime"
)

type Data struct {
	i, j int
}

func main() {
	var N = 10_000
	var structure []Data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, Data{value, value})
	}

	_ = structure[0]
	fmt.Println(structure)
	runtime.GC()
	fmt.Println("============== GC done")
}
