package main

import "testing"

var muTest = 0
var chTest = 0

/* func BenchmarkUseMutex(b *testing.B) {
	// run the Fib function b.N times
	var r int
	for n := 0; n < b.N; n++ {
		r = useMutex()
	}
	muTest = r
} */

func BenchmarkUseChannel(b *testing.B) {
	// run the Fib function b.N times
	var r int
	for n := 0; n < b.N; n++ {
		r = useChannnel()
	}
	chTest = r
}
