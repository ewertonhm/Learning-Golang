package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("Contents of: %v", fileName)

	} else {
		fmt.Printf("%v doesn't exist!", fileName)
	}

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	// setup a buffer
	buf := make([]byte, 1024)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			// nothing left
			break
		}

		// Check for other err
		if err != nil {
			fmt.Println(err)
		}
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
	}
}
