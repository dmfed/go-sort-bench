package main

import (
	"sort"
	"testing"
)

func BenchmarkIntSliceSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		s := getSliceCopy()
		b.StartTimer()
		sort.Sort(sort.IntSlice(s))
		b.StopTimer()
	}
}

func BenchmarkLessFuncSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		s := getSliceCopy()
		b.StartTimer()
		sort.Slice(s, func(a, b int) bool { return s[a] < s[b] })
		b.StopTimer()
	}
}

func BenchmarkInterfaceSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		var s custIntSlice
		s = getSliceCopy()
		b.StartTimer()
		sort.Sort(s)
		b.StopTimer()
	}
}
