// Garbage Collector
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var N = 1000
	myMap := make(map[int]int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[i] = value
	}

	runtime.GC()
	_ = myMap[0]
	fmt.Println(myMap)
}
