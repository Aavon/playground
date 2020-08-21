package blotdb

import (
	"fmt"
	"github.com/boltdb/bolt"
	"os"
	"testing"
	"time"
)

func Test_case1(t *testing.T) {
	db, err := bolt.Open("./test.db", os.ModePerm, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(db.GoString())
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("db0"))
		if err != nil {
			return err
		}
		bucket.Put([]byte("a"), []byte("1"))
		bucket.Put([]byte("b"), []byte("2"))
		bucket.Put([]byte("c"), []byte("3"))
		return nil
	})

	if err != nil {
		panic(err)
	}

	//db2, err := bolt.Open("./test.db", os.ModePerm, &bolt.Options{Timeout: 5 * time.Second})
	//if err != nil {
	//	panic(err)
	//}
	//defer db2.Close()
	//fmt.Println(db2.GoString())

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("db0"))
		if err != nil {
			return err
		}
		bucket.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})

	if err != nil {
		panic(err)
	}

}
