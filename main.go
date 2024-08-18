package main

import (
	"fmt"
	"sync"

	"github.com/utkarsh352/wallet-go/models"
	"github.com/utkarsh352/wallet-go/utility"
)

func main() {
	// Initialize the ledger and its mutex
	ledger := []models.Transaction{}
	ledgerMux := &sync.Mutex{}

	// Create wallets
	wallet1 := models.NewWallet()
	wallet1.Balance = 100

	wallet2 := models.NewWallet()
	wallet2.Balance = 50

	wallet3 := models.NewWallet()

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
			tx.From, tx.To, tx.Amount, utility.FormatTimestamp(tx.Timestamp))
	}
}
