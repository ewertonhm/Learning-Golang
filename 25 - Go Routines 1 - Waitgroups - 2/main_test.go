package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Hello, world!"

	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find 'Hello, world!', but it is not there")
	}
}

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("epsilon")
	wg.Wait()

	output := msg

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon, but it is not there")
	}
}

func Test_Main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected to find 'Hello, universe!', but it is not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expected to find 'Hello, cosmos!', but it is not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find 'Hello, world!', but it is not there")
	}
}
