package wallet

import (
	"testing"
)

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("Wanted an error but didn't get one")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Didn't wanted a error but got one")
	}
}

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

		err := wallet.Withdraw(10)
		assertNoError(t, err)

		got := wallet.Balance()
		want := Bitcoin(5)

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 5}

		err := wallet.Withdraw(Bitcoin(10))
		assertError(t, err, ErrInsufficientFunds)
	})

}
