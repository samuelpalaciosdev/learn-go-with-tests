package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat character 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectRepetition(t, got, expected)
	})
	t.Run("repeat character 3 times", func(t *testing.T) {
		got := Repeat("b", 3)
		expected := "bbb"
		assertCorrectRepetition(t, got, expected)
	})
	t.Run("repeat characters 0 times", func(t *testing.T) {
		got := Repeat("a", 0)
		expected := "a"
		assertCorrectRepetition(t, got, expected)
	})
}

func assertCorrectRepetition(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
