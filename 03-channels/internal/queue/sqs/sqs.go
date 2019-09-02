package sqs

import (
	"test/internal/queue"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"golang.org/x/xerrors"
)

// GetString returns a default value if the given string pointer is nil, otherwise returns the dereferenced string
// pointer.
func getString(strPtr *string, defaultValue string) string {
	if strPtr == nil {
		return defaultValue
	}
	return *strPtr
}

// SQS is a pretty thin abstraction of the AWS SQS library.
type SQS struct {
	*sqs.SQS
}

// New creates a new connection to AWS SQS.
func New(region, sqsEndpoint string) *SQS {

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("test", "test", ""),
	}))

	config := &aws.Config{Region: aws.String(region)}

	if sqsEndpoint != "" {
		config.Endpoint = aws.String(sqsEndpoint)
	}

	svc := sqs.New(sess, config)

	return &SQS{svc}
}

// Receive receives messages from an SQS queue.Message
func (s *SQS) Receive(queueURL string) ([]queue.Message, error) {

	returnMessages := []queue.Message{}
	maxNumberOfMessages := int64(10)

	result, err := s.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: &maxNumberOfMessages,
	})

	if err != nil {
		return returnMessages, xerrors.Errorf(`receiving message from queue "%s": %w`, queueURL, err)
	}

	for _, msg := range result.Messages {

		returnMessages = append(returnMessages, queue.Message{
			Body:          getString(msg.Body, ""),
			MessageID:     getString(msg.MessageId, ""),
			ReceiptHandle: getString(msg.ReceiptHandle, ""),
		})
	}

	return returnMessages, nil
}

// Delete deletes a message from a queue.
func (s *SQS) Delete(queueURL, receiptHandle string) error {
	_, err := s.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	})
	if err != nil {
		return xerrors.Errorf(`deleting SQS message "%s" for queue "%s": %w`, receiptHandle, queueURL, err)
	}
	return nil
}

// Send sends an SQS message.
func (s *SQS) Send(queueURL, message string) error {
	_, err := s.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(message),
	})

	if err != nil {
		return xerrors.Errorf("sending SQS message: %w", err)
	}
	return nil
}
