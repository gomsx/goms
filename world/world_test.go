package world

import "testing"

func TestWorld(t *testing.T) {
	var w World
	var want = "world\n"
	if got := w.Print(); got != want {
		t.Errorf("w.Print() = %v,want %v", got, want)
	}
}
