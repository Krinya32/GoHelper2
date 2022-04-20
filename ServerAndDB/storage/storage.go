package storage

import (
	"database/sql"
)

//Instance of storage
type Storage struct {
	config *Config
	//DataBAse FileDescriptor
	db *sql.DB
}

//Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

//Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return nil
	}
	storage.db = db
	return nil
}

//Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}
