package models

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type Wallet struct {
	Address uuid.UUID
	Balance float64
	mu      sync.Mutex // Mutex to make Wallet thread-safe
}

func NewWallet() *Wallet {
	address := uuid.New()
	return &Wallet{Address: address, Balance: 0}
}

func (w *Wallet) Send(amount float64, receiver *Wallet, ledger *[]Transaction, ledgerMux *sync.Mutex) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	receiver.mu.Lock()
	defer receiver.mu.Unlock()

	if w.Balance < amount {
		return errors.New("insufficient funds")
	}

	w.Balance -= amount
	receiver.Balance += amount

	RecordTransaction(w.Address, receiver.Address, amount, ledger, ledgerMux)

	return nil
}

func (w *Wallet) CheckBalance() float64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.Balance
}
