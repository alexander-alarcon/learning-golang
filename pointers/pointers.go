// Package pointers provides Wallet
package pointers

import "errors"

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrInvalidDeposit = errors.New("deposit amount must be greater than zero")

func (wallet *Wallet) Deposit(amount Bitcoin) error {
	if amount <= 0 {
		return ErrInvalidDeposit
	}

	wallet.balance += amount
	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsufficientFunds
	}

	wallet.balance -= amount
	return nil
}
