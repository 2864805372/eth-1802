package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	// 打开数据库
	db, err := bolt.Open("eg.db", 0600, nil)
	if nil != err {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		// 创建桶
		b, err := tx.CreateBucket([]byte("EgBucket"))
		if nil != err {
			return err
		}
		if nil != b {
			err = b.Put([]byte("1"),[]byte("110"))
			if nil != err {
				return err
			}
		}
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("EgBucket"))
		if nil != b {
			value := b.Get([]byte("1"))
			fmt.Printf("value : %s\n", string(value))
		}
		return nil
	})
}
