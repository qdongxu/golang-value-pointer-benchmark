package cmd

import (
	"fmt"
	"testing"
)

func BenchmarkCalcWithValue(b *testing.B) {
	var total = 0
	for i := 0; i < b.N; i++ {
		d := calcWithValue()
		total += sumWithValue(d)
	}
	fmt.Sprintf("%d", total)
}

func BenchmarkCalcWithPointer(b *testing.B) {
	var total = 0
	for i := 0; i < b.N; i++ {
		d := calcWithPointer()
		total += sumWithPointer(d)
	}
	fmt.Sprintf("%d", total)
}

func BenchmarkChanWithValue(b *testing.B) {
	var total = 0
	for i := 0; i < b.N; i++ {
		total += ChanWithValue()
	}
	fmt.Sprintf("%d", total)
}

func BenchmarkChanWithPointer(b *testing.B) {
	var total = 0
	for i := 0; i < b.N; i++ {
		total += ChanWithPointer()
	}
	fmt.Sprintf("%d", total)
}

func BenchmarkMethodValue(b *testing.B) {
	var total = 0
	for i := 0; i < b.N; i++ {
		d := NewDataValue()
		total += d.calcWithValue(10)
	}
	fmt.Sprintf("%d", total)
}

func BenchmarkMethodPointer(b *testing.B) {
	var total = 0
	for i := 0; i < b.N; i++ {
		d := NewDataPointer()
		total += d.calcWithPointer(10)
	}
	fmt.Sprintf("%d", total)
}
