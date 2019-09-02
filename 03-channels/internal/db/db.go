package db

import "fmt"

// Client is just a dummy. It does nothing.
type Client struct {
}

// New creates a new instance of the DB Client.
func New() *Client {
	return &Client{}
}

// InsertCreatedEvent inserts a "transaction created" event into the database.
func (c *Client) InsertCreatedEvent(event TxnCreatedEvent) error {
	fmt.Printf("inserting created event: %+v\n", event)
	return nil
}

// InsertFailedEvent inserts a "transaction failed" event into the database.
func (c *Client) InsertFailedEvent(event TxnFailedEvent) error {
	fmt.Printf("inserting failed event: %+v\n", event)
	return nil
}

// InsertCancelledEvent inserts a "transaction cancelled" event into the database.
func (c *Client) InsertCancelledEvent(event TxnCancelledEvent) error {
	fmt.Printf("inserting cancelled event: %+v\n", event)
	return nil
}
