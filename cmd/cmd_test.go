package cmd

import "testing"

func BenchmarkExecute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Execute()
	}
}
