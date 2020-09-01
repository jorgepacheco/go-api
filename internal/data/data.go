package data

import (
	"database/sql"
	"io/ioutil"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

// Data manages the connection to the database.
type Data struct {
	DB *sql.DB
}

// New returns a new instance of Data with the database connection ready.
func New() *Data {
	once.Do(initDB)

	return data
}

// initialize the data variable with the connection to the database.
func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	err = MakeMigration(db)
	if err != nil {
		log.Fatal("Error in db migration..", err)
		log.Panic(err)
	}

	data = &Data{
		DB: db,
	}
}

// MakeMigration creates all the tables in the database
func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}

// Close closes the resources used by data.
func Close() error {
	if data == nil {
		return nil
	}

	return data.DB.Close()
}
