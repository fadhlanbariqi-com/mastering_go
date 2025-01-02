// Subsection unsafe code
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Println(*pointer, "")
	memAddres := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memAddres))
		fmt.Println(*pointer, "")
		memAddres = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
}
