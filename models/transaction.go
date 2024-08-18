package models

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	From      uuid.UUID // Sender's wallet address
	To        uuid.UUID // Receiver's wallet address
	Amount    float64   // Amount transferred
	Timestamp time.Time // Time when the transaction was recorded
}

func RecordTransaction(from uuid.UUID, to uuid.UUID, amount float64, ledger *[]Transaction, ledgerMux *sync.Mutex) {
	ledgerMux.Lock()
	defer ledgerMux.Unlock()

	transaction := Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	*ledger = append(*ledger, transaction)
}
