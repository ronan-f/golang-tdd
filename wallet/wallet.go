package wallet

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
