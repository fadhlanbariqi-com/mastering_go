package main

// #cgo CFLAGS: -Ic_codes
// #cgo LDFLAGS: -Lc_codes -lcallC
// #include <stdlib.h>
// #include "callC.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C Function!")
	C.cHello()

	fmt.Println("Going to call another C Function!")
	myMessage := C.CString("This is Fadhlan B!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("All perfectly done!")
}
