package main

import "testing"

func TestUpper(t *testing.T) {
	result := Upper("hello")
	if result != "HELLO" {
		t.Fatalf("expected HELLO, got %s", result)
	}
}
