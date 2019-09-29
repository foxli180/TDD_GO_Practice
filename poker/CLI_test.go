package poker_test

import (
	"awesomeProject/poker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("Record Chris Win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("Record Cleo Win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
