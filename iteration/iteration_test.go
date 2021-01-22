package iteration

import "testing"

const N = 3

func TestRepeat(t *testing.T) {
	got := Repeat("a", N)
	want := "aaa"

	if got != want {
		t.Errorf("want '%s' but got '%s.\n", got, want)
	}
}

// Execute by <go test -bench=".">
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", N)
	}
}
