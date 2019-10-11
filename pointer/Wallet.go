package pointer

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

func (wallet *Wallet) Deposit(amount Bitcoin) {
	 wallet.balance += amount

}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

var InSufficientFundError = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) error{

	if amount > wallet.Balance() {
		return InSufficientFundError
	}
	wallet.balance -= amount
	return nil
}
