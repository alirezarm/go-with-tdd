package pointers

import (
	"errors"
	"fmt"
)


var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

// Implement the Stringer interface (from fmt package) for Bitcoin type 
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func(w *Wallet) Balance() Bitcoin {
	return w.balance
}

// error is an interface type
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}