// Garbage Collector

package main

import (
	"fmt"
	"runtime"
)

func main() {
	var N = 1000
	split := make([]map[int]int, 200)
	for i := range split {
		split[i] = make(map[int]int)
	}

	for i := 0; i < N; i++ {
		value := int(i)
		split[i%200][value] = value
	}

	runtime.GC()
	_ = split[0][0]
	fmt.Println(split)

}
