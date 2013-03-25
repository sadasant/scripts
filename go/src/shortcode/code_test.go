package shortcode

import (
	"testing"
)

// This test is weak
func Test_compare_two_codes6(t *testing.T) {
	code1 := Generate(6)
	code2 := Generate(6)
	println("Code 1:", code1)
	println("Code 2:", code2)
	if code1 == code2 {
		t.Errorf("This codes should be different", code1, code2)
	}
}

// Benchmarking code generation speed
func BenchmarkCodes6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(6)
	}
}
