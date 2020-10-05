package main

import "testing"

func BenchmarkNoLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoLock()
	}
}
