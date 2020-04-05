package service

import "testing"

func Test_getUid(t *testing.T) {
	initUidGenerator()
	for i := 0; i < 100; i++ {
		setUidMax(50)
		t.Run("getUid()", func(t *testing.T) {
			if got := getUid(); got < 0 || got > getUidMax() {
				t.Errorf("getUid() = %v, want > %v && < %v", got, 0, getUidMax())
			}
		})
	}
}
