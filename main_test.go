package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Chris")
	expected := "Hello, Chris!"

	if result != expected {
		t.Errorf("Hello() = %q; expected %q", result, expected)
	}
}
