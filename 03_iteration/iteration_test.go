package iteration

import "testing"
import "fmt"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	str := Repeat("a", 3)
	fmt.Println(str)
	// Output: aaa
}

func TestRepeat(t *testing.T) {
	t.Run("4 repeat", func(t *testing.T) {
		repeated := Repeat("a", 4)
		expected := "aaaa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("5 repeat", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func assert(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}