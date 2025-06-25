package main

import "testing"

func BenchmarkHttpGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ReusableFetch("https://httpbin.org/get")
		if err != nil {
			b.Fatal(err)
		}

	}
}

// go test -bench=. -benchmem
