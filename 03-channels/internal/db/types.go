package db

import "time"

// TxnCreatedEvent describes a "transaction created" event.
type TxnCreatedEvent struct {
	ID          string    `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	Amount      int       `json:"amount"`
}

// TxnFailedEvent describes a "transaction failed" event.
type TxnFailedEvent struct {
	ID         string    `json:"id"`
	FailedDate time.Time `json:"failedDate"`
	ReasonCode int       `json:"reasonCode"`
}

// TxnCancelledEvent describes a "transaction cancelled" event.
type TxnCancelledEvent struct {
	ID            string    `json:"id"`
	CancelledDate time.Time `json:"cancelledDate"`
}
