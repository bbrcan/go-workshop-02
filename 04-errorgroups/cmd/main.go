package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"test/internal/db"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	// set global random number generator seed
	rand.Seed(time.Now().Unix())
	dbClient := db.New()

	result := struct {
		Store     db.Store
		Locations []db.Location
		MenuItems []db.MenuItem
	}{}

	// Use an ErrorGroup so we can run our queries concurrently, and also handle any errors

	group := errgroup.Group{}

	group.Go(func() error {
		store, err := dbClient.GetStoreInfo()
		if err == nil && store != nil {
			result.Store = *store
		}
		return err
	})

	group.Go(func() error {
		locations, err := dbClient.GetStoreLocations()
		if err == nil {
			result.Locations = locations
		}
		return err
	})

	group.Go(func() error {
		menuItems, err := dbClient.GetMenuItems()
		if err == nil {
			result.MenuItems = menuItems
		}
		return err
	})

	// Wait for our tasks to finish

	if err := group.Wait(); err != nil {
		log.Fatalf("Error fetching store data: %v", err)
	}

	// Convert result to JSON and print

	b, err := json.MarshalIndent(result, "", "  ")

	if err != nil {
		log.Fatalf("Error marshalling result to json: %v", err)
	}

	fmt.Println(string(b))
}
