package main

import (
	"fmt"
	"net/http"
)

func setName(name string) {
	fmt.Println("Hello", name)
}

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}

	var my1 = struct {
		Name string
		Age  int
	}{
		"John Doe",
		22,
	}

	http.HandleFunc("/", handlerIndex)

	fmt.Println(my1)

	setName(*&my1.Name)
}
