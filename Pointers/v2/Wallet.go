package main

import "fmt"

type Stringer interface {
	String() string
}
type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
func (w *Wallet) Ballance() Bitcoin {
	return w.balance
}
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
