package main

import (
    "sync"
    "testing"
)

func BenchmarkSyncPool(b *testing.B) {
	bytesPool := &sync.Pool{New: func() interface{} { return make([]byte, 0, 200) }}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts := bytesPool.Get().([]byte)
		bytesPool.Put(bts)
	}
}

func BenchmarkMake(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
        _ = make([]byte, 0, 200)
	}
}
