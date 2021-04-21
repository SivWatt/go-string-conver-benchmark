package demo

import (
	"testing"
)

func BenchmarkIntStringFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intStringFmt(i)
	}
}

func BenchmarkIntStringItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intStringItoa(i)
	}
}

func BenchmarkIntStringFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intStringFormatInt(int64(i))
	}
}
