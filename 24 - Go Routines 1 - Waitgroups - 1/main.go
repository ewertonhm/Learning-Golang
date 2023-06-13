package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	// defer: when this is done
	// Wg.done: decrements the wg by 1
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	// create the waitgroup
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	// add integer to wait groups (how many times it needs to wait)
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	// wait the wg reachs zero
	wg.Wait()
}
