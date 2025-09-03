package pointers

import (
	"LearningGo/testhelpers"
	"testing"
)

func AssertWalletBalance(t *testing.T, wallet *Wallet, want Bitcoin, testName string) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("For test case %q: got %v, want %v", testName, got, want)
	}
}

func TestWallet(t *testing.T) {
	tests := []testhelpers.StateTestCaseNoExpected[*Wallet]{
		{
			Name: "initial_state_zero_balance",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(wallet *Wallet) {},
			Assert: func(t *testing.T, wallet *Wallet, testName string) {
				AssertWalletBalance(t, wallet, 0.0, testName)
			},
		},
		{
			Name: "deposit_once",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(wallet *Wallet) {
				wallet.Deposit(0.5)
			},
			Assert: func(t *testing.T, wallet *Wallet, name string) {
				AssertWalletBalance(t, wallet, 0.5, name)
			},
		},
		{
			Name: "deposit_accumulated",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(wallet *Wallet) {
				wallet.Deposit(0.3)
				wallet.Deposit(0.7)
			},
			Assert: func(t *testing.T, wallet *Wallet, name string) {
				AssertWalletBalance(t, wallet, 1.0, name)
			},
		},
	}

	testhelpers.RunStateTestsNoExpected(t, tests)
}
