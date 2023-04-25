package main

import "testing"

func Example_Main() {
	goMain([]string{})
	// Output:
	// Hello World
}

func Test_Main(t *testing.T) {
	if stasus := goMain([]string{}); stasus != 0 {
		t.Error("Expected 0, got ", stasus)
	}
}
