package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := strings.Repeat("a", 6)
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}

}
func ExampleRepeat() {
	repeated := Repeat("w", 8)
	fmt.Println(repeated)
	// Output: wwwwwwww
}
