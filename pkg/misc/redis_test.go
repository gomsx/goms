package misc

import "testing"

func TestGetRedisKey(t *testing.T) {
	type args struct {
		uid int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "any1", args: args{uid: 88}, want: "uid#88"},
		{name: "any2", args: args{uid: 99}, want: "uid#99"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisKey(tt.args.uid); got != tt.want {
				t.Errorf("GetRedisKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
