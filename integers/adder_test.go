package integers

import (
	"testing"
)

func TestAdder(t *testing.T)  {

	assertExpectedSum := func(t *testing.T, sum, expected int) {
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("check 2 + 2 equls 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertExpectedSum(t, sum, expected)
	})

	t.Run("Another add check if it hard coded", func(t *testing.T) {
		sum := Add(1, 5)
		expected := 6
		assertExpectedSum(t, sum, expected)
	})

}
