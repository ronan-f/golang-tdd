package wallet

import "fmt"

// Wallet is a struct for storing currency
type Wallet struct {
	balance Bitcoin
}

// Bitcoin type
type Bitcoin int

// Deposit adds to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance shows current wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw deducts from wallet balance
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
