package wallet

import "fmt"

type Bitcoin float64

type Wallet struct{
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
} 

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
