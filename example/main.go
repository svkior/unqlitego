package main

import (
	"fmt"
	"log"

	"github.com/svkior/unqlitego"
)

func main() {
	// New
	db, err := unqlitego.NewDatabase("/root/unqlite.db")
	if err != nil {
		log.Fatalf("store error: %v", err)
	}
	defer db.Close()

	k := []byte(`k1`)
	v := []byte(`v1`)

	// Store
	err = db.Store(k, v)
	if err != nil {
		log.Fatalf("store error: %v", err)
	}

	// Query
	bytes, err := db.Fetch(k)
	if err != nil {
		log.Fatalf("fetch error: %v", err)
	}
	fmt.Printf("key %s, value %s\n", string(k), string(bytes))

	// Update
	v2 := []byte(`中文`)
	err = db.Store(k, v2)
	if err != nil {
		log.Fatalf("update error: %v", err)
	}

	// Query again
	bytes, err = db.Fetch(k)
	if err != nil {
		log.Fatalf("fetch error: %v", err)
	}
	fmt.Printf("key %s,value %s\n", string(k), string(bytes))

	// Delete
	err = db.Delete(k)
	if err != nil {
		log.Fatalf("delete error: %v", err)
	}

	// Query again
	bytes, err = db.Fetch(k)
	if err != nil {
		// No such record
		log.Printf("fetch error: %v", err)
	}
	//fmt.Printf("key %s,value %s\n", string(k), string(bytes))

	// Store more data
	db.Store([]byte(`k5`), []byte(`v5`))
	db.Store([]byte(`k7`), []byte(`v7`))
	db.Store([]byte(`k6`), []byte(`v6`))
	db.Store([]byte(`k2`), []byte(`v2`))
	db.Store([]byte(`k3`), []byte(`v3`))
	db.Store([]byte(`k4`), []byte(`v4`))

	// Commit changes to disk
	err = db.Commit()
	if err != nil {
		log.Fatalf("commit error: %v", err)
	}

	// Cursor
	cursor, err := db.NewCursor()
	if err != nil {
		log.Fatalf("new cursor error: %v", err)
	}

	// match exactly
	key := []byte(`k6`)
	err = cursor.Seek(key)
	if err != nil {
		log.Fatalf("seek error: %v", err)
	}
	bytes, err = cursor.Value()
	if err != nil {
		log.Fatalf("get cursor value error: %v", err)
	}
	fmt.Printf("key %s, value %s\n", key, bytes)

	// get next
	err = cursor.Next()
	if err != nil {
		log.Fatalf("point to next error: %v", err)
	}
	key, err = cursor.Key()
	if err != nil {
		log.Fatalf("get cursor key error: %v", err)
	}
	bytes, err = cursor.Value()
	if err != nil {
		log.Fatalf("get cursor value error: %v", err)
	}
	fmt.Printf("key %s, value %s\n", key, bytes)

	// get first
	err = cursor.First()
	if err != nil {
		log.Fatalf("point to next error: %v", err)
	}
	key, err = cursor.Key()
	if err != nil {
		log.Fatalf("get cursor key error: %v", err)
	}
	bytes, err = cursor.Value()
	if err != nil {
		log.Fatalf("get cursor value error: %v", err)
	}
	fmt.Printf("key %s, value %s\n", key, bytes)
}

// How to run this example?
// Follow REAME.md in https://github.com/zyfdegh/unqlite-crud

// More examples?
// Combine these pages to get better understanding.
// Imported unqlite.go. (https://github.com/svkior/unqlitego/blob/master/unqlite.go)
// UnQLite C/C++ API Reference - List Of Functions. (https://unqlite.org/c_api_func.html)
