package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	addr := ":" + os.Args[2]
	fs := http.FileServer(http.Dir(httpDir))
	fmt.Printf("Server starting at 0.0.0.0%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, fs))
}
