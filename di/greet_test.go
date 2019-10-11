package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T)  {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Fox")

	got := buffer.String()
	want := "Hello Fox"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
