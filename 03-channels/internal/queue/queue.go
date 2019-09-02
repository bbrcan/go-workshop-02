package queue

// Message represents a message received from SQS.
type Message struct {
	// A unique identifier for the message.
	MessageID string
	// An identifier associated with the act of receiving the message. A new receipt handle is returned every time you
	// receive a message.
	ReceiptHandle string
	// The message's contents (non URL-encoded).
	Body string
}

// Receiver receives messages from SQS.
type Receiver interface {
	// Receive messages from the queue.
	Receive(queueURL string) ([]Message, error)
	// Delete message from the queue.
	Delete(queueURL, receiptHandle string) error
}
