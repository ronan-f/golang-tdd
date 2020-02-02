package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("Wanted an error but didn't get one")
		}
	}

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.balance

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}

		wallet.Withdraw(Bitcoin(3))

		want := Bitcoin(2)

		assertBalance(t, wallet, want)

	})

	t.Run("Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(2)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})
}
