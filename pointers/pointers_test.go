package pointers

import (
	"LearningGo/testhelpers"
	"testing"
)

func TestWallet(t *testing.T) {
	tests := []testhelpers.StateTestCase[*Wallet, float64]{
		{
			Name: "initial_state_zero_balance",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(wallet *Wallet) float64 {
				return wallet.Balance()
			},
			Expected: 0.0,
			Assert:   testhelpers.AssertFloatEqual,
		},
		{
			Name: "deposit_once",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(wallet *Wallet) float64 {
				wallet.Deposit(0.5)
				return wallet.Balance()
			},
			Expected: 0.5,
			Assert:   testhelpers.AssertFloatEqual,
		},
		{
			Name: "deposit_accumulated",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(wallet *Wallet) float64 {
				wallet.Deposit(0.3)
				wallet.Deposit(0.7)
				return wallet.Balance()
			},
			Expected: 1.0,
			Assert:   testhelpers.AssertFloatEqual,
		},
	}

	testhelpers.RunStateTests[*Wallet, float64](t, tests)
}
