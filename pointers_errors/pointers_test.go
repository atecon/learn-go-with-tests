package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Test Deposit()", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Test Withdraw() with nil error", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, nil)
	})

	t.Run("Test Withdraw() with error", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s but want %s", got.String(), want.String())
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if want != nil && got != want {
		t.Errorf("wanted error '%q' but got '%q'", got, want)
	} else if want == nil && got != nil {
		t.Errorf("wanted no error but got '%q'.", got)
	}
}
