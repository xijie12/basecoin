package main

import (
	"gocode/20190724/go-bili/basecoin/bolt"
	"fmt"
	"log"
)

func main(){

	db, err := bolt.Open("D:/lubo/gowork/src/gocode/20190724/go-bili/basecoin/test.db",0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		}
		bucket.Put([]byte("1"),[]byte("hello"))
		bucket.Put([]byte("2"),[]byte("world"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("b1"))
		if b == nil {
			log.Panic("bucket b1 没有")
		}
		v := b.Get([]byte("1"))
		v1 := b.Get([]byte("2"))
		fmt.Printf("The answer is: %s\n", v)
		fmt.Printf("The answer is: %s\n", v1)
		return nil
	})
}
