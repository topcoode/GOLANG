package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5433
	user     = "postgres"
	password = "mahi"
	dbname   = "postgres"
)

var names = []string{
	"Alice", "Bob", "Charlie", "David", "Emma", "Frank", "Grace", "Henry", "Ivy", "Jack",
	"Katie", "Liam", "Mia", "Noah", "Olivia", "Patrick", "Quinn", "Rachel", "Sam", "Tina",
	"Uma", "Victor", "Wendy", "Xander", "Yara", "Zoe",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Connect to the database
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Insert random names
	for i := 0; i < 50; i++ {
		name := names[rand.Intn(len(names))] // Randomly select a name from the list
		age := rand.Intn(80) + 18            // Random age between 18 and 97
		err := insertName(db, name, age)
		if err != nil {
			log.Println("Error inserting name:", err)
		}
	}

	fmt.Println("Insertion of random names completed.")
}

func insertName(db *sql.DB, name string, age int) error {
	query := "INSERT INTO subscribers (name, age) VALUES ($1, $2)"
	_, err := db.Exec(query, name, age)
	return err
}
