package model

import (
	"testing"
)

func TestGetUid(t *testing.T) {
	for i := 0; i < 100; i++ {
		SetUidMax(50)
		t.Run("getUid()", func(t *testing.T) {
			if got := GetUid(); got < 0 || got > GetUidMax() {
				t.Errorf("GetUid() = %v, want > %v && < %v", got, 0, GetUidMax())
			}
		})
	}
}
