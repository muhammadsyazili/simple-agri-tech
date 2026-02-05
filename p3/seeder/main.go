package main

import (
	"coding-test/p3/config"
	"coding-test/p3/models"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"gorm.io/gorm/clause"
)

func main() {
	// Initialize DB connection
	config.ConnectDatabase()

	// Seed Users
	seedUsers()

	// Seed Spendings
	seedSpendings()

	// Reset Sequences
	resetSequences()
}

func resetSequences() {
	var result int
	config.DB.Raw("SELECT setval(pg_get_serial_sequence('users', 'id'), COALESCE(MAX(id), 1)) FROM users").Scan(&result)
	config.DB.Raw("SELECT setval(pg_get_serial_sequence('spendings', 'id'), COALESCE(MAX(id), 1)) FROM spendings").Scan(&result)
	fmt.Println("Sequences reset successfully")
}

func seedUsers() {
	file, err := os.Open("p3/seeder/users.tsv")
	if err != nil {
		log.Fatalf("Unable to read input file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	_, err = reader.Read()
	if err != nil {
		log.Fatalf("Failed to read header: %v", err)
	}

	var users []models.User

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading record: %v", err)
			continue
		}

		id, _ := strconv.Atoi(record[0])

		user := models.User{
			ID:               uint(id),
			Country:          record[1],
			CreditCardType:   record[2],
			CreditCardNumber: record[3],
			FirstName:        record[4],
			LastName:         record[5],
		}
		users = append(users, user)
	}

	batchSize := 100
	if err := config.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(users, batchSize).Error; err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}

	fmt.Printf("Successfully seeded %d users\n", len(users))
}

func seedSpendings() {
	file, err := os.Open("p3/seeder/spendings.tsv")
	if err != nil {
		log.Fatalf("Unable to read input file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	_, err = reader.Read()
	if err != nil {
		log.Fatalf("Failed to read header: %v", err)
	}

	var spendings []models.Spending

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading record: %v", err)
			continue
		}

		id, _ := strconv.Atoi(record[0])
		userID, _ := strconv.Atoi(record[1])
		totalBuy, _ := strconv.Atoi(record[2])

		spending := models.Spending{
			ID:       uint(id),
			UserID:   uint(userID),
			TotalBuy: int64(totalBuy),
		}
		spendings = append(spendings, spending)
	}

	batchSize := 100
	if err := config.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(spendings, batchSize).Error; err != nil {
		log.Fatalf("Failed to seed spendings: %v", err)
	}

	fmt.Printf("Successfully seeded %d spendings\n", len(spendings))
}
