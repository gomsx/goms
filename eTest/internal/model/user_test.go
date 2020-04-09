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

func TestCheckUid(t *testing.T) {
	var min int64 = 0
	var max int64 = 0xFFFF_FFFF
	SetUidMax(max)

	type args struct {
		uid int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "for succ min",
			args: args{
				uid: min,
			},
			want: true,
		},
		{
			name: "for succ min+1",
			args: args{
				uid: min + 1,
			},
			want: true,
		},
		{
			name: "for succ max",
			args: args{
				uid: max,
			},
			want: true,
		},
		{
			name: "for succ max-1",
			args: args{
				uid: max - 1,
			},
			want: true,
		},
		{
			name: "for succ rand",
			args: args{
				uid: GetUid(),
			},
			want: true,
		},
		{
			name: "for failed min-1",
			args: args{
				uid: min - 1,
			},
			want: false,
		},
		{
			name: "for failed max+1",
			args: args{
				uid: max + 1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUid(tt.args.uid); got != tt.want {
				t.Errorf("CheckUid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckUidS(t *testing.T) {
	var max int64 = 12345678
	SetUidMax(max)
	type args struct {
		uidstr string
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
	}{
		{
			name: "for succ 12345",
			args: args{
				uidstr: "12345",
			},
			want:  12345,
			want1: true,
		},
		{
			name: "for failed -12345",
			args: args{
				uidstr: "-12345",
			},
			want:  -12345,
			want1: false,
		},
		{
			name: "for failed > max",
			args: args{
				uidstr: "123456789",
			},
			want:  123456789,
			want1: false,
		},
		{
			name: "for failed 0xFFFF",
			args: args{
				uidstr: "0xFFFF",
			},
			want:  -1,
			want1: false,
		},
		{
			name: "for failed 123AAAA",
			args: args{
				uidstr: "123AAAA",
			},
			want:  -1,
			want1: false,
		},
		{
			name: "for failed AAAA",
			args: args{
				uidstr: "AAAA",
			},
			want:  -1,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckUidS(tt.args.uidstr)
			if got != tt.want {
				t.Errorf("CheckUidS() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckUidS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
