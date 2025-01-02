package main

import (
	"errors"
	"fmt"
)

func testError(a, b int) error {
	if a == b {
		err := errors.New("a dan b sama")
		return err
	} else {
		return nil
	}
}

func main() {
	err := testError(1, 2)
	if err == nil {
		fmt.Println("testError() ended normally!")
	} else {
		fmt.Println(err)
	}
	err = testError(10, 10)
	if err == nil {
		fmt.Println("testError() ended normally!")
	} else {
		fmt.Println(err)
	}
	if err.Error() == "Error in testError() function!" {
		fmt.Println("!!")
	}
}
