package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("uknown word", func(t *testing.T) {
		_, got := dictionary.Search("uknown")
		assertErrors(t, got, ErrNotFound)
	})
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
