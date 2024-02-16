package main

import "fmt"

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}
func (w Wallet) Ballance() int {
	return w.balance
}
