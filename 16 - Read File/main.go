package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Reading File:", os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(readFile(file))
}

func readFile(f *os.File) string {
	fb := make([]byte, 32*1024)
	f.Read(fb)
	return string(fb)
}
