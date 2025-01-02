package main

import (
	"fmt"
	"log"
	"os"
)

const pathLog = "/tmp/customlog.log"

func main() {
	f, err := os.OpenFile(pathLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "customlog ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags | log.Llongfile)
	iLog.Println("Hello ini custom log from Go")
	iLog.Println("Hari jumat")
	//
}
