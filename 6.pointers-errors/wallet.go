package wallet

import "fmt"

type Wallet struct {
	balance int
}

// * Using pointer to not use a copy of wallet, but instead change the original values within it
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	return w.balance
}
