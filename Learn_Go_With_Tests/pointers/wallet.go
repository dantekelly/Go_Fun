package main

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(a Bitcoin) {
	w.balance += a
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) Withdraw(a Bitcoin) error {
	if a > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= a
	return nil
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
