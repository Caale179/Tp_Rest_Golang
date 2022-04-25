package persistance

import (
	"github.com/boltdb/bolt"
	"log"
)

type BoltDb struct {
	Db *bolt.DB
}

func (d *BoltDb) DbOpen(fileName string) {
	var err error
	d.Db, err = bolt.Open(fileName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Db *bolt.DB) {
		err := Db.Close()
		if err != nil {

		}
	}(d.Db)
}

func (d *BoltDb) DbClose() {
	err := d.Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *BoltDb) doPath() string {
	return d.Db.Path()
}

func (d *BoltDb) DbCreateBucket(name string) {
	err := d.Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
}

func (d *BoltDb) DbPut(bucketName string, key string, value string) {
	err := d.Db.Update(func(tx *bolt.Tx) error {
		d.DbCreateBucket(bucketName)
		bucket := tx.Bucket([]byte(bucketName))
		err := bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

}

func (d *BoltDb) dbGet(bucketName string, key string) string {
	var value []byte
	err := d.Db.Update(func(tx *bolt.Tx) error {
		d.DbCreateBucket(bucketName)
		bucket := tx.Bucket([]byte(bucketName))
		value = bucket.Get([]byte(key))
		return nil
	})
	if err != nil {
		return ""
	}
	return string(value)
}

func (d *BoltDb) dbGetAll(bucketName string) []string {
	var res []string
	err := d.Db.Update(func(tx *bolt.Tx) error {
		d.DbCreateBucket(bucketName)
		bucket := tx.Bucket([]byte(bucketName))
		err := bucket.ForEach(func(k, v []byte) error {
			res = append(res, string(k))
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return res
}

func (d *BoltDb) dbDelete(bucketName string, key string) {
	err := d.Db.Update(func(tx *bolt.Tx) error {
		d.DbCreateBucket(bucketName)
		bucket := tx.Bucket([]byte(bucketName))
		err := bucket.Delete([]byte(key))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
}
