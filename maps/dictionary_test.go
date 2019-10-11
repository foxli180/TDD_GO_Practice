package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(got, want, t)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrorNotFound

		if err == nil {
			t.Fatal("expected get an error.")
		}
		assertError(err, want, t)
	})
}

func TestAdd(t *testing.T) {

	t.Run("Add a new word with definition", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		assertDefinition(dictionary, word, definition, t)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(err, ErrorWordExists, t)
		assertDefinition(dictionary, word, definition, t)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		dictionary.Update(word, newDefinition)
		assertDefinition(dictionary, word, newDefinition, t)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(err, ErrorWordNotExist, t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(err, ErrorNotFound, t)
	})

	t.Run("not existing", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(err, ErrorNotFound, t)
	})

}

func assertDefinition(dictionary Dictionary, word, definition string, t *testing.T) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}
	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}

func assertStrings(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(got error, want error, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
