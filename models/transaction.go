package models

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	From      uuid.UUID 
	To        uuid.UUID 
	Amount    float64   
	Timestamp time.Time 
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
