package model

import (
	"testing"
)

func TestGetUid(t *testing.T) {
	SetUidMax(5)
	for i := 0; i < 10; i++ {
		t.Run("getUid()", func(t *testing.T) {
			if got := GetUid(); got < 0 || got > GetUidMax() {
				t.Errorf("GetUid() = %v, want > %v && < %v", got, 0, GetUidMax())
			}
		})
	}
}
