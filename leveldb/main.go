package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	KEY = "TESTKEY"
)

func main() {
	db, err := leveldb.OpenFile("/tmp/data/db2", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Put([]byte(KEY), []byte("This is test value."), nil); err != nil {
		panic(err)
	}
	data, err := db.Get([]byte(KEY), nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %v, value: %v", KEY, string(data))

	if err = db.Delete([]byte(KEY), nil); err != nil {
		panic(err)
	}
}
