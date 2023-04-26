package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("Insufficient funds in your balance")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	fmt.Printf("address of balance in Withdraw is %p \n", &w.balance)

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}
