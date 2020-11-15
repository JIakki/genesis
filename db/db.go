package db

import (
	"github.com/boltdb/bolt"
	"log"
)

type DB struct {
	*bolt.DB
}

func New() *DB {
	connection, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal("DB error: ", err)
	}

	db := &DB{connection}

	return db
}
