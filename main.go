package main

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Address uuid.UUID
	Balance float64
	mu      sync.Mutex // Mutex to make Wallet thread-safe
}

type Transaction struct {
	From      uuid.UUID
	To        uuid.UUID
	Amount    float64
	Timestamp time.Time
}

func NewWallet() *Wallet {
	address := uuid.New()
	return &Wallet{Address: address, Balance: 0}
}

func (w *Wallet) Send(amount float64, receiver *Wallet, ledger *[]Transaction, ledgerMux *sync.Mutex) error {
	// Lock the sender wallet
	w.mu.Lock()
	defer w.mu.Unlock()

	// Lock the receiver wallet
	receiver.mu.Lock()
	defer receiver.mu.Unlock()

	if w.Balance < amount {
		return errors.New("insufficient funds")
	}

	// Perform the transaction
	w.Balance -= amount
	receiver.Balance += amount

	// Record the transaction in the ledger using RecordTransaction
	RecordTransaction(w.Address, receiver.Address, amount, ledger, ledgerMux)

	return nil
}

func (w *Wallet) CheckBalance() float64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.Balance
}

func RecordTransaction(from uuid.UUID, to uuid.UUID, amount float64, ledger *[]Transaction, ledgerMux *sync.Mutex) {
	// Lock the ledger to ensure thread-safety
	ledgerMux.Lock()
	defer ledgerMux.Unlock()

	// Create a new transaction
	transaction := Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	*ledger = append(*ledger, transaction)
}

func FormatTimestamp(t time.Time) string {
	return t.Format("15:04:05 02/01/2006")
}

func main() {
	// Initialize the ledger and its mutex
	ledger := []Transaction{}
	ledgerMux := &sync.Mutex{}

	// Create wallets
	wallet1 := NewWallet()
	wallet1.Balance = 100

	wallet2 := NewWallet()
	wallet2.Balance = 50

	wallet3 := NewWallet()

	// Perform transactions
	err := wallet1.Send(30, wallet2, &ledger, ledgerMux)
	if err != nil {
		fmt.Println("Transaction 1 failed:", err)
	} else {
		fmt.Println("Transaction 1 succeeded")
	}

	err = wallet2.Send(20, wallet3, &ledger, ledgerMux)
	if err != nil {
		fmt.Println("Transaction 2 failed:", err)
	} else {
		fmt.Println("Transaction 2 succeeded")
	}

	err = wallet3.Send(10, wallet1, &ledger, ledgerMux)
	if err != nil {
		fmt.Println("Transaction 3 failed:", err)
	} else {
		fmt.Println("Transaction 3 succeeded")
	}

	// Check balances
	fmt.Println("Wallet 1 balance:", wallet1.CheckBalance())
	fmt.Println("Wallet 2 balance:", wallet2.CheckBalance())
	fmt.Println("Wallet 3 balance:", wallet3.CheckBalance())

	// Print ledger
	fmt.Println("\nTransaction Ledger:")
	for _, tx := range ledger {
		fmt.Printf("From: %s, To: %s, Amount: %.2f, Timestamp: %s\n",
			tx.From, tx.To, tx.Amount, FormatTimestamp(tx.Timestamp))
	}
}
