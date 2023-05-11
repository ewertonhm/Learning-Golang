package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "ewerton" {
		return "$1$5vQo5DS8$Lje.C.H27.wKcTey7gMDu1"
	}
	return ""
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	addr := ":" + os.Args[2]

	authenticator := auth.NewBasicAuthenticator("files.lab", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))

	fmt.Printf("Server starting at 0.0.0.0%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
