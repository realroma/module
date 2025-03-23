package tests

import "testing"

func TestRightSum(t *testing.T) {
	got := sumRight(1, 2)
	want := 3
	if got != want {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}

// Нажать серые надписи над функцией для тестов.
func TestWrongSum(t *testing.T) {
	got := sumWrong(1, 2)
	want := 3
	if got != want {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}
