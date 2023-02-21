package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "https://puresearch-production.up.railway.app",
		APIKey: os.Getenv("ADMIN_API_KEY"),
	})

	connStr := os.Getenv("DATABASE_URL")
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query database

	// Users
	users, err := db.Query("SELECT uuid, name, email, handle FROM users")
	if err != nil {
		log.Fatal(err)
	}
	var userDocs []interface{}
	for users.Next() {
		var myuuid uuid.UUID
		var name string
		var email string
		var handle string

		err = users.Scan(&myuuid, &name, &email, &handle)
		if err != nil {
			log.Fatal(err)
		}
		userDocs = append(userDocs, map[string]interface{}{
			"uuid":   myuuid,
			"name":   name,
			"email":  email,
			"handle": handle,
		})
	}

	// Organizations
	orgs, err := db.Query("SELECT uuid, name, handle, description FROM organizations")
	if err != nil {
		log.Fatal(err)
	}
	var orgDocs []interface{}
	for orgs.Next() {
		var myuuid uuid.UUID
		var name string
		var handle string
		var description string

		err = orgs.Scan(&myuuid, &name, &handle, &description)
		if err != nil {
			log.Fatal(err)
		}
		orgDocs = append(orgDocs, map[string]interface{}{
			"uuid":        myuuid,
			"name":        name,
			"handle":      handle,
			"description": description,
		})
	}

	// Models
	models, err := db.Query("SELECT uuid, name, wiki, organization_uuid, is_public FROM models")
	if err != nil {
		log.Fatal(err)
	}
	var modelDocs []interface{}
	for models.Next() {
		var myuuid uuid.UUID
		var name string
		var wiki string
		var organizationUuid uuid.UUID
		var is_public bool

		err = models.Scan(&myuuid, &name, &wiki, &organizationUuid, &is_public)
		if err != nil {
			log.Fatal(err)
		}
		modelDocs = append(modelDocs, map[string]interface{}{
			"uuid":              myuuid,
			"name":              name,
			"wiki":              wiki,
			"organization_uuid": organizationUuid,
			"is_public":         is_public,
		})
	}

	// Datasets
	datasets, err := db.Query("SELECT uuid, name, wiki, organization_uuid, is_public FROM datasets")
	if err != nil {
		log.Fatal(err)
	}
	var datasetDocs []interface{}
	for datasets.Next() {
		var myuuid uuid.UUID
		var name string
		var wiki string
		var organizationUuid uuid.UUID
		var is_public bool

		err = datasets.Scan(&myuuid, &name, &wiki, &organizationUuid, &is_public)
		if err != nil {
			log.Fatal(err)
		}
		datasetDocs = append(datasetDocs, map[string]interface{}{
			"uuid":              myuuid,
			"name":              name,
			"wiki":              wiki,
			"organization_uuid": organizationUuid,
			"is_public":         is_public,
		})
	}

	// Add documents to index
	_, err = client.Index("users").AddDocuments(userDocs, "uuid")
	if err != nil {
		panic(err)
	}
	_, err = client.Index("organizations").AddDocuments(orgDocs, "uuid")
	if err != nil {
		panic(err)
	}
	_, err = client.Index("models").AddDocuments(modelDocs, "uuid")
	if err != nil {
		panic(err)
	}
	_, err = client.Index("datasets").AddDocuments(modelDocs, "uuid")
	if err != nil {
		panic(err)
	}
}
