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
			Test: func(t *testing.T, wallet *Wallet, name string) {},
			Assert: func(t *testing.T, wallet *Wallet, testName string) {
				AssertWalletBalance(t, wallet, 0.0, testName)
			},
		},
		{
			Name: "deposit_once",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(t *testing.T, wallet *Wallet, name string) {
				err := wallet.Deposit(0.5)
				testhelpers.AssertNoError(t, err, name)
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
			Test: func(t *testing.T, wallet *Wallet, name string) {
				err := wallet.Deposit(0.5)
				testhelpers.AssertNoError(t, err, name)

				err = wallet.Deposit(0.5)
				testhelpers.AssertNoError(t, err, name)
			},
			Assert: func(t *testing.T, wallet *Wallet, name string) {
				AssertWalletBalance(t, wallet, 1.0, name)
			},
		},
		{
			Name: "deposit_negative_amount",
			Setup: func() *Wallet {
				return &Wallet{}
			},
			Test: func(t *testing.T, wallet *Wallet, name string) {
				err := wallet.Deposit(-1.0)
				testhelpers.AssertError(t, err, "deposit amount must be greater than zero", name)
			},
			Assert: func(t *testing.T, wallet *Wallet, name string) {
				AssertWalletBalance(t, wallet, 0.0, name)
			},
		},
		{
			Name: "withdraw_once",
			Setup: func() *Wallet {
				return &Wallet{balance: 1.0}
			},
			Test: func(t *testing.T, wallet *Wallet, name string) {
				err := wallet.Withdraw(0.5)
				testhelpers.AssertNoError(t, err, name)
			},
			Assert: func(t *testing.T, wallet *Wallet, name string) {
				AssertWalletBalance(t, wallet, 0.5, name)
			},
		},
		{
			Name: "withdraw_to_zero",
			Setup: func() *Wallet {
				return &Wallet{balance: 1.0}
			},
			Test: func(t *testing.T, wallet *Wallet, name string) {
				err := wallet.Withdraw(1.0)
				testhelpers.AssertNoError(t, err, name)
			},
			Assert: func(t *testing.T, wallet *Wallet, name string) {
				AssertWalletBalance(t, wallet, 0.0, name)
			},
		},
		{
			Name: "withdraw_more_than_balance",
			Setup: func() *Wallet {
				return &Wallet{balance: 1.0}
			},
			Test: func(t *testing.T, wallet *Wallet, name string) {
				err := wallet.Withdraw(1.5)
				testhelpers.AssertError(t, err, "cannot withdraw, insufficient funds", name)
			},
			Assert: func(t *testing.T, wallet *Wallet, name string) {
				AssertWalletBalance(t, wallet, 1.0, name)
			},
		},
		{
			Name: "withdraw_all_then_deposit",
			Setup: func() *Wallet {
				return &Wallet{balance: 1.0}
			},
			Test: func(t *testing.T, wallet *Wallet, name string) {
				err := wallet.Withdraw(1.0)
				testhelpers.AssertNoError(t, err, name)

				err = wallet.Deposit(2.0)
				testhelpers.AssertNoError(t, err, name)
			},
			Assert: func(t *testing.T, wallet *Wallet, name string) {
				AssertWalletBalance(t, wallet, 2.0, name)
			},
		},
	}

	testhelpers.RunStateTestsNoExpected(t, tests)
}
