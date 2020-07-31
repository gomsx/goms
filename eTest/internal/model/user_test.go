package model

import (
	"testing"
)

func TestGetRedisKey(t *testing.T) {
	type args struct {
		uid int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "xx",
			args: args{
				uid: 88,
			},
			want: "uid#88",
		}, {
			name: "yy",
			args: args{
				uid: 99,
			},
			want: "uid#99",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisKey(tt.args.uid); got != tt.want {
				t.Errorf("GetRedisKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
