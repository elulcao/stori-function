package processor

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MSQL struct {
	DB *sql.DB
}

// openDB opens a database connection to the specified DSN.
func openDB(dsn string) (msqldb *MSQL, err error) {
	// Open a database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	//Create a new SDB object with the opened database
	msqldb = &MSQL{
		DB: db,
	}

	return msqldb, nil
}

// createTable creates a new table in the database.
func (m *MSQL) createTable() (err error) {
	// Create the transactions table if it doesn't exist
	_, err = m.DB.Exec("CREATE TABLE IF NOT EXISTS transactions (id INT NOT NULL AUTO_INCREMENT, user VARCHAR(255), month VARCHAR(255), transaction DECIMAL(10,5), PRIMARY KEY (id))")
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

// insert inserts a new record into the database.
func (m *MSQL) insert(txn *Txn) (err error) {
	// Insert a new transaction record
	_, err = m.DB.Exec("INSERT INTO transactions (user, month, transaction) VALUES (?, ?, ?)", txn.User, txn.Date, txn.Txn)
	if err != nil {
		return fmt.Errorf("failed to insert record: %w", err)
	}

	return nil
}

// Insert inserts a new record into the database.
func Insert(txn *Txn) (err error) {
	db, err := openDB("user:password@tcp(172.20.0.4:3306)/demo_db")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.DB.Close()

	err = db.createTable()
	if err != nil {
		return fmt.Errorf("create (if not exist) action failed: %w", err)
	}

	err = db.insert(txn)
	if err != nil {
		return fmt.Errorf("insert action failed: %w", err)
	}

	return nil
}
