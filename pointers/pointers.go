// Package pointers provides Wallet
package pointers

type Wallet struct {
	balance float64
}

func (wallet *Wallet) Deposit(amount float64) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() float64 {
	return wallet.balance
}
