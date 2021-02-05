package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dict.Search("foo")

		assertError(t, err, ErrNotFound)

	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q but want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q, given %s", got, want, "test")
	}
}
