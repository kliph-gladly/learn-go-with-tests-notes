package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(tb testing.TB, wallet Wallet, want Bitcoin) {
		tb.Helper()

		got := wallet.Balance()

		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		if got != want {
			tb.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(tb testing.TB, got error, want error) {
		tb.Helper()

		if got == nil {
			tb.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			tb.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(tb testing.TB, got error) {
		tb.Helper()

		if got != nil {
			tb.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
