// Package pointers provides Wallet
package pointers

import "errors"

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}

	wallet.balance -= amount
	return nil
}
