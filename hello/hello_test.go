package main

import "fmt"
import "testing"

func TestHello(t *testing.T) {
	NAME := "Chris"
	got := Hello(NAME)
	want := fmt.Sprintf("Hello, %s", NAME)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Subtests
func TestHelloSub(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// tell the test suite that this method is a helper
		// when it fails the line number reported will be in our function
		// call rather than inside our test helper.
		t.Helper()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Empty string supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
