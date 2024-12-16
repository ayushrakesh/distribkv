package db

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

var defaultBucket = []byte("default")
var replicaBucket = []byte("replication")

// database is an open bolt database
type Database struct {
	db *bolt.DB
}

func (d *Database) createBuckets() error {
	return d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(defaultBucket)
		return err
	})
}

// new database returns new instance of a database that we can work with
func NewDatabase(dbpath string) (db *Database, closeFunc func() error, err error) {
	boltDB, err := bolt.Open(dbpath, 0600, nil)
	if err != nil {
		return nil, nil, err
	}
	closeFunc = boltDB.Close
	db = &Database{
		db: boltDB,
	}

	if err := db.createBuckets(); err != nil {
		closeFunc()
		return nil, nil, fmt.Errorf("creating default bucket: %w", err)
	}

	return db, closeFunc, nil
}

func (db *Database) SetKey(key string, value []byte) error {
	return db.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(defaultBucket).Put([]byte(key), value)
	})
}

func (db *Database) GetKey(key string) ([]byte, error) {
	var result []byte
	err := db.db.View(func(tx *bolt.Tx) error {
		result = tx.Bucket(defaultBucket).Get([]byte(key))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
