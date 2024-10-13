package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Block struct {
	BtcHeight  int64
	BtcHash    string
	Timestamp  time.Time
	Coinbase   string
	EvmAddress string
	EvmHeight  int64
}

func main() {
	// Secure SQLite connection
	db, err := sql.Open("sqlite3", "./block.db?_pragma=journal_mode(WAL)&_pragma=foreign_keys(ON)")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table and indexes
	createTableSQL := `CREATE TABLE IF NOT EXISTS Block (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		btcHeight INTEGER,
		btcHash TEXT,
		timestamp DATETIME,
		coinbase TEXT,
		evmAddress TEXT,
		evmHeight INTEGER
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createIndexSQL := `CREATE INDEX IF NOT EXISTS idx_timestamp ON Block (timestamp);
					   CREATE INDEX IF NOT EXISTS idx_evmHeight ON Block (evmHeight);`
	_, err = db.Exec(createIndexSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Insert sample data
	block := Block{
		BtcHeight:  123456,
		BtcHash:    "0000000000000000000762d6b883d6bd8",
		Timestamp:  time.Now(),
		Coinbase:   "coinbase-sample",
		EvmAddress: "0xabcdef123456789",
		EvmHeight:  987654,
	}
	err = insertBlock(db, block)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Block inserted successfully")

	// Query by timestamp and evmHeight
	timestamp := block.Timestamp
	evmHeight := block.EvmHeight
	results, err := queryByTimestampAndEvmHeight(db, timestamp, evmHeight)
	if err != nil {
		log.Fatal(err)
	}

	// Output results
	for _, result := range results {
		fmt.Printf("Block: %+v\n", result)
	}
}

func insertBlock(db *sql.DB, block Block) error {
	insertSQL := `INSERT INTO Block (btcHeight, btcHash, timestamp, coinbase, evmAddress, evmHeight) 
				  VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(insertSQL, block.BtcHeight, block.BtcHash, block.Timestamp, block.Coinbase, block.EvmAddress, block.EvmHeight)
	if err != nil {
		return err
	}
	return nil
}

func queryByTimestampAndEvmHeight(db *sql.DB, timestamp time.Time, evmHeight int64) ([]Block, error) {
	querySQL := `SELECT btcHeight, btcHash, timestamp, coinbase, evmAddress, evmHeight
				 FROM Block 
				 WHERE timestamp = ? AND evmHeight = ?`

	rows, err := db.Query(querySQL, timestamp, evmHeight)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blocks []Block
	for rows.Next() {
		var block Block
		err = rows.Scan(&block.BtcHeight, &block.BtcHash, &block.Timestamp, &block.Coinbase, &block.EvmAddress, &block.EvmHeight)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}
