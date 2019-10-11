package hello

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Say hello with default ''", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("say hello with people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Fox", "Spanish")
		want := "Hola Fox"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("in franch", func(t *testing.T) {
		got := Hello("Fox", "French")
		want := "Bonjour Fox"
		assertCorrectMessage(t, got, want)
	})

}

