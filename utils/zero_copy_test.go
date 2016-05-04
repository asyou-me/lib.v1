package utils

import (
	"strings"
	"testing"
)

var s1 = strings.Repeat("a", 1024)
var s2 = strings.Repeat("a", 1024)

func test() {
	b := []byte(s1)
	_ = string(b)
}

func test2() {
	b := StrToBytes(s2)
	_ = BytesToStr(b)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test2()
	}
}
