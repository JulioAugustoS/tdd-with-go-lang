package main

import "testing"

func TestHello(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result %q; expected %q", result, expected)
		}
	}


	t.Run("saying hello to people", func(t *testing.T) {
		result := Hello("Chris")
		expected := "Hello, Chris!"

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World!"

		verifyCorrectMessage(t, result, expected)
	})
}
