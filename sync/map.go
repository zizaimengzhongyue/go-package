package main

import (
	"fmt"
	"sync"
)

var m sync.Map = sync.Map{}

func Store() {
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")
	m.Store("key4", "value4")
	val, ok := m.Load("key1")
	fmt.Printf("Sync.Map.Store: %+v, %t\n", val, ok)
}

func Delete() {
	m.Delete("key1")
	val, ok := m.Load("key1")
	fmt.Printf("Sync.Map.Delete: %+v, %t\n", val, ok)
}

func LoadOrStore() {
	val, ok := m.LoadOrStore("key1", "value1")
	fmt.Printf("Sycn.Map.LoadOrStore: %+v, %t\n", val, ok)
	val, ok = m.LoadOrStore("key2", "value3")
	fmt.Printf("Sync.Map.LoadOrStore: %+v, %t\n", val, ok)
	val, ok = m.Load("key2")
	fmt.Printf("Sync.Map.Load: %+v, %t\n", val, ok)
}

func Load() {
	val, ok := m.Load("key1")
	fmt.Printf("Sync.Map.Load: %+v, %t\n", val, ok)
	val, ok = m.Load("key5")
	fmt.Printf("Sync.Map.Load: %+v, %t\n", val, ok)
}

func Range() {
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Sync.Map.Range: %+v, %+v\n", key, value)
		return true
	})
}

func main() {
	Store()
	Delete()
	LoadOrStore()
	Load()
	Range()
}
