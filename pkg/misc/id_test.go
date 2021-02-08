package misc

import (
	"testing"
)

func TestGenUid(t *testing.T) {
	uidx := GenUid()
	uidy := GenUid()
	uidz := GenUid()
	tests := []struct {
		name    string
		compute bool
		want    bool
	}{
		{"succx1", uidx-0 > 0, true},
		{"succx2", uidy-uidx > 0, true},
		{"succx3", uidz-uidy > 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.compute; got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenRequestId(t *testing.T) {
	ridx := GenRequestId()
	ridy := GenRequestId()
	ridz := GenRequestId()
	tests := []struct {
		name    string
		compute bool
		want    bool
	}{
		{"succx1", ridx-0 > 0, true},
		{"succx2", ridy-ridx > 0, true},
		{"succx3", ridz-ridy > 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.compute; got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
