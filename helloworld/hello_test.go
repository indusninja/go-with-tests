package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("say 'Hello, World' when an empty string is supplied, in english", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to Chris, in english", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to Christine, in english", func(t *testing.T) {
		got := Hello("Christine", "english")
		want := "Hello, Christine"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied, in spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to Chris, in spanish", func(t *testing.T) {
		got := Hello("Chris", "spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to Christine, in spanish", func(t *testing.T) {
		got := Hello("Christine", "spanish")
		want := "Hola, Christine"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied, in french", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to Chris, in french", func(t *testing.T) {
		got := Hello("Chris", "french")
		want := "Bonjour, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to Christine, in french", func(t *testing.T) {
		got := Hello("Christine", "french")
		want := "Bonjour, Christine"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
