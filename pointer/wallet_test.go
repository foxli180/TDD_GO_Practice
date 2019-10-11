package pointer

import "testing"

func TestWallet(t *testing.T)  {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %d want %d", wallet.Balance(), want)
		}
	}

	assertNoError := func(got error, t *testing.T) {
		if got != nil {
			t.Fatal("got an error but did not want one")
		}
	}

	assertError := func(got error, want error, t *testing.T) {
		if got == nil {
			t.Fatal("wanted an error but did not get one")
		}
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(err, t)
	})
	
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		staringBalance := Bitcoin(20)
		wallet := Wallet{staringBalance}

		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, staringBalance)
		assertError(err, InSufficientFundError, t)
	})
}
