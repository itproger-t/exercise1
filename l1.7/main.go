package main

import (
	"fmt"
	"sync"
)

type DataStore struct {
	mt    sync.Mutex
	cache map[string]string
}

func NewDataStore() *DataStore {
	return &DataStore{
		cache: make(map[string]string),
	}
}

func (ds *DataStore) set(key string, value string) {
	ds.cache[key] = value
}

func (ds *DataStore) get(key string) string {
	if ds.count() > 0 {
		return ds.cache[key]
	}
	return ""
}

func (ds *DataStore) count() int {
	return len(ds.cache)
}

func (ds *DataStore) Set(key string, value string) {
	ds.mt.Lock()
	defer ds.mt.Unlock()
	ds.set(key, value)
}

func (ds *DataStore) Get(key string) string {
	ds.mt.Lock()
	defer ds.mt.Unlock()
	return ds.get(key)
}

func (ds *DataStore) Count() int {
	ds.mt.Lock()
	defer ds.mt.Unlock()
	return ds.count()
}

func main() {
	store := NewDataStore()
	store.Set("Go", "Lang")
	result := store.Get("Go")
	fmt.Println(result)
}
