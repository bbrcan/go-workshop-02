package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	ID    string    `json:"id"`
	Date  time.Time `json:"date"`
	Count int       `json:"count"`
}

func main() {

	// Convert from JSON string to struct

	messageJSON := `{ "id": "1234", "date": "2019-03-26T04:16:46Z", "count": 1 }`
	message := Message{}

	if err := json.Unmarshal([]byte(messageJSON), &message); err != nil {
		fmt.Println("Unmarshalling json:", err)
		return
	}

	fmt.Printf("Message: %+v\n", message)

	// Convert from struct to JSON

	b, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Marshalling json:", err)
		return
	}

	fmt.Printf("JSON: %s\n", string(b))
}
