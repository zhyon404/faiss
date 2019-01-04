package main

import (
	"testing"
)

func Benchmark_Search(b *testing.B) {
	d, nb, xb, nq, xq := data()
	idx := index(d)
	add(idx, nb, xb)
	b.ResetTimer()
	//b.RunParallel(func(pb *testing.PB) {
	//	for pb.Next() {
	//		searching(idx, nq, xq, 4)
	//	}
	//})
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			searching(idx, nq, xq, 4)
		}
	}
}
