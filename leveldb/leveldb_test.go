package leveldb

import (
	"fmt"
	"github.com/golang/leveldb"
	//"github.com/golang/leveldb/db"
	"testing"
)

func Test_case(t *testing.T) {
	db, err := leveldb.Open("test.db", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Set([]byte("a"), []byte("1"), nil)
	if err != nil {
		panic(err)
	}
	data, err := db.Get([]byte("a"), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
