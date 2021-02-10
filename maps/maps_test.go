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

func TestAdd(t *testing.T) {

	t.Run("add new key-value pair", func(t *testing.T) {
		dict := Dictionary{}
		// var dict = map[string]string{} // won't work as Add() requires type Dictionary
		key := "test"
		value := "added valued"
		dict.Add(key, value)

		want := "added valued"
		got, error := dict.Search("test")

		assertError(t, error, nil)
		assertString(t, got, want)
	})

	t.Run("add already existing key-value pair", func(t *testing.T) {
		key := "test"
		dict := Dictionary{key: "added valued"}

		newValue := "new value"
		dict.AddWithCheck(key, newValue)

		want := "added valued"
		got, error := dict.Search("test")

		assertError(t, error, nil)
		assertString(t, got, want)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update...", func(t *testing.T) {
		key := "test"
		value := "original value"
		dict := Dictionary{key: value}
		newValue := "new value"

		dict.Update(key, newValue)
		assertDefinition(t, dict, key, newValue)

	})

}

func TestDelete(t *testing.T) {

	t.Run("delete", func(t *testing.T) {
		key := "test"
		dict := Dictionary{key: "test def"}

		dict.Delete(key)
		_, err := dict.Search(key)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", key)
		}
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

func assertDefinition(t testing.TB, dict Dictionary, key, newValue string) {
	t.Helper()
	if dict[key] != newValue {
		t.Errorf("Got %s but want %s.", dict[key], newValue)
	}
}
