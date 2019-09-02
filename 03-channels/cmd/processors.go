package main

import (
	"encoding/json"
	"test/internal/db"
	"test/internal/queue"

	"golang.org/x/xerrors"
)

func processTxnCreatedMessage(dbClient *db.Client, msg queue.Message) error {
	event := db.TxnCreatedEvent{}
	if err := json.Unmarshal([]byte(msg.Body), &event); err != nil {
		return xerrors.Errorf("unmarshalling json: %w", err)
	}
	if err := dbClient.InsertCreatedEvent(event); err != nil {
		return xerrors.Errorf("inserting event into db: %w", err)
	}
	return nil
}

func processTxnFailedMessage(dbClient *db.Client, msg queue.Message) error {
	event := db.TxnFailedEvent{}
	if err := json.Unmarshal([]byte(msg.Body), &event); err != nil {
		return xerrors.Errorf("unmarshalling json: %w", err)
	}
	if err := dbClient.InsertFailedEvent(event); err != nil {
		return xerrors.Errorf("inserting event into db: %w", err)
	}
	return nil
}

func processTxnCancelledMessage(dbClient *db.Client, msg queue.Message) error {
	event := db.TxnCancelledEvent{}
	if err := json.Unmarshal([]byte(msg.Body), &event); err != nil {
		return xerrors.Errorf("unmarshalling json: %w", err)
	}
	if err := dbClient.InsertCancelledEvent(event); err != nil {
		return xerrors.Errorf("inserting event into db: %w", err)
	}
	return nil
}
