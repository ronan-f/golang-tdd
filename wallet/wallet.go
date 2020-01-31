package wallet

// Wallet is a struct for storing currency
type Wallet struct {
	balance int
}

// Deposit adds to wallet
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance shows current wallet balance
func (w *Wallet) Balance() int {
	return w.balance
}
