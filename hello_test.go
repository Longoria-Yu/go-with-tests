package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Longoria")
		expected := "Hello, Longoria"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, World"
		assertCorrectMessage(t, actual, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	/*
		t.Helper() is needed to tell the test suite that this method is a helper.
		By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		This will help other developers track down problems easier.
	*/
	t.Helper()
	if actual != expected {
		t.Errorf("got %q but want %q", actual, expected)
	}
}
