package main

import (
	"encoding/json"
	"math/rand"
	"test/internal/db"
	"test/internal/queue/sqs"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func sendJSON(sqsClient *sqs.SQS, queueURL string, payload interface{}) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return xerrors.Errorf("marshalling payload to json: %w", err)
	}
	return sqsClient.Send(queueURL, string(b))
}

func randomName(randGen *rand.Rand) string {
	options := []string{"Bob Dylan", "Leonard Cohen", "Miles Davis", "Janis Joplin", "Nina Simone", "Neil Young",
		"Neil Diamond", "John Coltrane", "Ella Fitzgerald", "Sarah Vaughan", "John Lennon", "Joni Mitchell", "Paul Simon",
		"Cat Stevens", "Nick Drake", "Nick Cave"}
	// create random index
	idx := randGen.Intn(len(options))
	return options[idx]
}

func main() {

	sqsClient := sqs.New("ap-southeast-2", "http://localhost:4576")

	txnCreatedQueueURL := "http://aws:4576/queue/txn-created-queue"
	txnFailedQueueURL := "http://aws:4576/queue/txn-failed-queue"
	txnCancelledQueueURL := "http://aws:4576/queue/txn-cancelled-queue"

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	sendJSON(sqsClient, txnCreatedQueueURL, db.TxnCreatedEvent{
		ID:          uuid.New().String(),
		CreatedDate: time.Now(),
		From:        randomName(randGen),
		To:          randomName(randGen),
		Amount:      randGen.Intn(500),
	})

	sendJSON(sqsClient, txnFailedQueueURL, db.TxnFailedEvent{
		ID:         uuid.New().String(),
		FailedDate: time.Now(),
		ReasonCode: randGen.Intn(10),
	})

	sendJSON(sqsClient, txnCancelledQueueURL, db.TxnCancelledEvent{
		ID:            uuid.New().String(),
		CancelledDate: time.Now(),
	})
}
