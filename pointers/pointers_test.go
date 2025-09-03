package pointers

import (
	"LearningGo/testhelpers"
	"testing"
)

func AssertBitcoinEqual(t *testing.T, got, want Bitcoin, testName string) {
	t.Helper()
	if got != want {
		t.Errorf("For test case %q: got %v, want %v", testName, got, want)
	}
}

func TestWallet(t *testing.T) {
	tests := []testhelpers.StateTestCase[*Wallet, Bitcoin]{
		{
			Name: "initial_state_zero_balance",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(wallet *Wallet) Bitcoin {
				return wallet.Balance()
			},
			Expected: 0.0,
			Assert:   AssertBitcoinEqual,
		},
		{
			Name: "deposit_once",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(wallet *Wallet) Bitcoin {
				wallet.Deposit(0.5)
				return wallet.Balance()
			},
			Expected: 0.5,
			Assert:   AssertBitcoinEqual,
		},
		{
			Name: "deposit_accumulated",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(wallet *Wallet) Bitcoin {
				wallet.Deposit(0.3)
				wallet.Deposit(0.7)
				return wallet.Balance()
			},
			Expected: 1.0,
			Assert:   AssertBitcoinEqual,
		},
	}

	testhelpers.RunStateTests(t, tests)
}
