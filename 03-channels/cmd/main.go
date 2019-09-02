package main

import (
	"log"
	"test/internal/db"
	"test/internal/queue"
	"test/internal/queue/sqs"
)

func receive(receiver queue.Receiver, queueURL string, msgChan chan<- queue.Message, errChan chan<- error) {
	for {
		messages, err := receiver.Receive(queueURL)

		if err != nil {
			errChan <- err
			return
		}
		for _, msg := range messages {
			msgChan <- msg
		}
	}
}

func main() {

	txnCreatedQueueURL := "http://aws:4576/queue/txn-created-queue"
	txnFailedQueueURL := "http://aws:4576/queue/txn-failed-queue"
	txnCancelledQueueURL := "http://aws:4576/queue/txn-cancelled-queue"

	sqsClient := sqs.New("ap-southeast-2", "http://localhost:4576")
	dbClient := db.New()

	txnCreatedMessages := make(chan queue.Message, 5)
	txnFailedMessages := make(chan queue.Message, 5)
	txnCancelledMessages := make(chan queue.Message, 5)
	errChan := make(chan error, 5)

	go receive(sqsClient, txnCreatedQueueURL, txnCreatedMessages, errChan)
	go receive(sqsClient, txnFailedQueueURL, txnFailedMessages, errChan)
	go receive(sqsClient, txnCancelledQueueURL, txnCancelledMessages, errChan)

	for {
		select {
		case msg := <-txnCreatedMessages:
			if err := processTxnCreatedMessage(dbClient, msg); err != nil {
				log.Printf("failed to process txn created message: %v", err)
			}
			sqsClient.Delete(txnCreatedQueueURL, msg.ReceiptHandle)
		case msg := <-txnFailedMessages:
			if err := processTxnFailedMessage(dbClient, msg); err != nil {
				log.Printf("failed to process txn failed message: %v", err)
			}
			sqsClient.Delete(txnFailedQueueURL, msg.ReceiptHandle)
		case msg := <-txnCancelledMessages:
			if err := processTxnCancelledMessage(dbClient, msg); err != nil {
				log.Printf("failed to process txn cancelled message: %v", err)
			}
			sqsClient.Delete(txnCancelledQueueURL, msg.ReceiptHandle)
		case err := <-errChan:
			log.Printf("error processing sqs message: %v", err)
		}
	}
}
