package database

import (
	"github.com/boltdb/bolt"
)

var BucketName = []byte("paths")

type Database struct {
	db *bolt.DB
}

func New(path string) (*Database, error) {
	db, err := bolt.Open(path, 0600, nil)

	if err != nil {
		return nil, err
	}

	return &Database{
		db,
	}, nil
}

func (d *Database) FindPath(path string) ([]byte, error) {
	var alias []byte

	err := d.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(BucketName)
		if err != nil {
			return err
		}

		curs := b.Cursor()

		for {
			k, v := curs.Next()

			if k == nil || v == nil {
				return ErrPathNotFound
			}

			if string(k) == path {
				alias = v
			}
		}
	})

	if err != nil {
		return alias, err
	}

	return alias, nil
}

func (d *Database) List() (map[string]string, error) {
	var paths = make(map[string]string, 10)

	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(BucketName)
		curs := b.Cursor()

		for k, v := curs.First(); k != nil; k, v = curs.Next() {
			paths[string(k)] = string(v)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}
