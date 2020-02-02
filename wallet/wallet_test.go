package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(5)}

		err := wallet.Withdraw(Bitcoin(3))

		want := Bitcoin(2)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(2)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrorInsufFunds)
	})
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("Got an error but didn't want one")
	}
}
