package db

import (
	"math/rand"
	"time"
)

// Client holds a connection to a database.
type Client struct {
}

// New creates a new Client instance.
func New() *Client {
	return &Client{}
}

// GetStoreLocations returns all the store locations.
func (c *Client) GetStoreLocations() ([]Location, error) {

	// Sleep for a random amount of time to simulate a DB query.
	randNum := rand.Float64()
	time.Sleep(time.Duration(randNum) * time.Second)

	return []Location{
		{
			ID:       "1234",
			Address:  "1 Test Way",
			Postcode: 3333,
			State:    "VIC",
			Country:  "Australia",
		},
		{
			ID:       "1235",
			Address:  "2 Test Way",
			Postcode: 3331,
			State:    "VIC",
			Country:  "Australia",
		},
		{
			ID:       "1236",
			Address:  "3 Test Way",
			Postcode: 3332,
			State:    "VIC",
			Country:  "Australia",
		},
	}, nil
}

// GetStoreInfo returns store information.
func (c *Client) GetStoreInfo() (*Store, error) {

	// Sleep for a random amount of time to simulate a DB query.
	randNum := rand.Float64()
	time.Sleep(time.Duration(randNum) * time.Second)

	return &Store{
		ID:   "1111",
		Name: "McDonald's",
	}, nil
}

// GetMenuItems returns all the store menu items.
func (c *Client) GetMenuItems() ([]MenuItem, error) {

	// Sleep for a random amount of time to simulate a DB query.
	randNum := rand.Float64()
	time.Sleep(time.Duration(randNum) * time.Second)

	return []MenuItem{
		{
			ID:       "1111",
			Name:     "Big Mac",
			Calories: 2000,
		},
		{
			ID:       "11112",
			Name:     "Cheeseburger",
			Calories: 1000,
		},
		{
			ID:       "11113",
			Name:     "Hamburger",
			Calories: 9000,
		},
	}, nil
}
