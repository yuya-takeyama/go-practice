package main

import "testing"

func BenchmarkUseIncrementOperator(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useIncrementOperator()
	}
}

func BenchmarkUseAtomicAddUint64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useAtomicAddUint32()
	}
}

func BenchmarkUseSyncMutexLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useSyncMutexLock()
	}
}
