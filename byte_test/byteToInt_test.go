package main

import "testing"

// from fib_test.go
func BenchmarkByteToInt(b *testing.B) {
	// run the Fib function b.N times
	test_byte := make([]byte, 2)
	test_byte[0] = 255
	test_byte[1] = 45
	for n := 0; n < b.N; n++ {
		byteToInt(test_byte)
	}
}

// from fib_test.go
func BenchmarkByteDecode(b *testing.B) {
	// run the Fib function b.N times
	test_byte := make([]byte, 2)
	test_byte[0] = 255
	test_byte[1] = 45
	for n := 0; n < b.N; n++ {
		byteDecode(test_byte)
	}
}
