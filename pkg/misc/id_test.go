package misc

import (
	"testing"
)

func TestGenUid(t *testing.T) {
	uidfoo := GenUid()
	uidbar := GenUid()
	tests := []struct {
		name    string
		compute bool
		want    bool
	}{
		{"Id should be lager than zero", uidfoo-0 > 0, true},
		{"Id should be ascending", uidbar-uidfoo > 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.compute; got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
