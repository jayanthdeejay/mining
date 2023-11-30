package store

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/jayanthdeejay/mining/config"
)

func DbOpen() (*sql.DB, string) {
	// Read the config file
	file, err := os.Open("../config/store.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := config.StoreConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	// Initialize database connection
	uriString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)
	db, err := sql.Open("postgres", uriString)
	if err != nil {
		fmt.Println("Error opening database:", err)
		panic(err)
	}

	row := db.QueryRow("SELECT pk FROM state ORDER BY timestamp DESC LIMIT 1")
	initial := ""
	err = row.Scan(&initial)
	if err == sql.ErrNoRows {
		fmt.Println("Error reading table:", err)
		return db, "0000000000000000000000000000000000000000000000000000000000000000"
	} else if err != nil {
		fmt.Println("Error while reading from state table:", err)
		panic(err)
	} else {
		return db, initial
	}
}

func Save(db *sql.DB, pk string) error {
	_, err := db.Exec("INSERT INTO state (pk, timestamp) VALUES ($1, $2)", pk, time.Now())
	return err
}
