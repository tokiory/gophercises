package main

import (
	"flag"
	"strings"
	"url-shortener/database"

	"github.com/boltdb/bolt"
)

var pathsToUrls = map[string]string{
	"/hello": "https://nohello.net/en/",
	"/arch":  "https://archive.org/details/texts",
}

var flagPath = flag.String("f", "urls.db", "path to database file")

func main() {
	flag.Parse()

	dbPath := *flagPath
	if !strings.HasSuffix(dbPath, ".db") {
		dbPath += ".db"
	}

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		if b := tx.Bucket(database.BucketName); b != nil {
			err := tx.DeleteBucket(database.BucketName)
			if err != nil {
				return err
			}
		}

		b, err := tx.CreateBucket(database.BucketName)
		if err != nil {
			return err
		}

		for k, v := range pathsToUrls {
			err := b.Put([]byte(k), []byte(v))
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
