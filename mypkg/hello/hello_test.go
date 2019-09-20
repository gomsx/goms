package hello

import "testing"

func TestHello(t *testing.T) {
	var h Hello
	var want = "hello\n"
	if got := h.Print(); got != want {
		t.Errorf("h.Print() = %v,want %v", got, want)
	}
}
