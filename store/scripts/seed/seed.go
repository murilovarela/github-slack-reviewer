package main

import (
	"log"

	"encore.app/store"
)

func main() {
	// Create a new store service
	dsn := "host=localhost user=testuser password=PGU2yYtqZ+pyJraDtdj2Tkb3GgW4KcqT dbname=github_slack_bot port=5432 sslmode=disable"
	storeSvc, err := store.NewStoreService(&dsn)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	// Create a new organization
	organization := store.Organization{
		Name:  "MyOrganization",
		Email: "myorganization@email.com",
	}

	// Create a new config
	config := store.Config{
		TimeToSummary:  60,
		TimeToReminder: 30,
	}

	storeSvc.CreateOrganization(organization, config)
}
