package part6

import (
	"fmt"
	"errors"
)

var ErrInsufficientFunds = "cannot withdraw, insufficient funds"

type Bitcoin int

type Wallet struct{
 balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
 w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}
func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return errors.New(ErrInsufficientFunds)

    }

	w.balance = w.balance - amount
	return nil
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}