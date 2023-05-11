package main

import (
	"os"
	"testing"
)

func BenchmarkReaddir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Open a new file handle for every loop iteration to avoid caching
		f, err := os.Open("/proc")
		if err != nil {
			b.Error(err)
		}

		_, err = f.Readdir(0)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkReadDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Open a new file handle for every loop iteration to avoid caching
		f, err := os.Open("/proc")
		if err != nil {
			b.Error(err)
		}

		_, err = f.ReadDir(0)
		if err != nil {
			b.Error(err)
		}
	}
}
