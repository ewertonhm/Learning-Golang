package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	path := "C:\\Users\\132160\\git\\Learning-Golang"
	addr := ":80"
	fs := http.FileServer(http.Dir(path))
	fmt.Println("Server starting at 0.0.0.0", addr)
	log.Fatal(http.ListenAndServe(addr, fs))
}
