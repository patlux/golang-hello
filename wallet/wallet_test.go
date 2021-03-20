package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := Bitcoin(11)

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}