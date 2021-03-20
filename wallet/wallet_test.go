package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 15}

		wallet.Withdraw(10)

		got := wallet.Balance()
		want := Bitcoin(5)

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}

	})

}
