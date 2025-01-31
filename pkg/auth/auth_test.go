package auth

import (
	"testing"
)

func TestAuthenticate(t *testing.T) {
	authenticated := Authenticate("user", "test")
	if !authenticated {
		t.Errorf("Should have successfully authenticated but got false")
	}
}

func BenchmarkIntMin(b *testing.B) {
	result := 1
	for i := 0; i < b.N; i++ {
		result += i
	}
}
