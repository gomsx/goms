package dao

import (
	"context"
	"reflect"
	"testing"

	m "github.com/fuwensun/goms/eApi/internal/model"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
)

func TestExistUserCC(t *testing.T) {
	ctx := context.Background()
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	cc, err := redis.Dial("tcp", s.Addr())

	Convey("Set a user to redis", t, func() {
		adao := &dao{redis: cc}
		user := m.GetUser()
		key := getRedisKey(user.Uid)
		cc.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...)

		Convey("When check this user from redis", func() {
			exist, err := adao.existUserCC(ctx, user.Uid)

			Convey("Then the result is exist", func() {
				So(err, ShouldBeNil)
				So(exist, ShouldBeTrue)
			})
		})

		Convey("When check other user from redis", func() {
			userx := m.GetUser()
			exist, err := adao.existUserCC(ctx, userx.Uid)

			Convey("Then the result is not exist", func() {
				So(err, ShouldBeNil)
				So(exist, ShouldBeFalse)
			})
		})

		Convey("When close connect, check this user from redis", func() {
			cc.Close()
			exist, err := adao.existUserCC(ctx, user.Uid)

			Convey("Then the result is err", func() {
				So(err, ShouldNotBeNil)
				So(exist, ShouldBeFalse)
			})
		})
	})
}

func TestSetUserCC(t *testing.T) {
	ctx := context.Background()
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	cc, err := redis.Dial("tcp", s.Addr())

	Convey("Given a user data", t, func() {
		adao := &dao{redis: cc}
		user := m.GetUser()

		Convey("When set this user to redis", func() {
			err := adao.setUserCC(ctx, user)

			Convey("Then the result is succ", func() {
				So(err, ShouldBeNil)

				Convey("When set same user to redis", func() {
					err := adao.setUserCC(ctx, user)

					Convey("Then the result is succ", func() {
						So(err, ShouldBeNil)
					})
				})
			})
		})

		Convey("When set other user to redis", func() {
			userx := m.GetUser()
			err := adao.setUserCC(ctx, userx)

			Convey("Then the result is succ", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When close connect, set this user from redis", func() {
			cc.Close()
			err := adao.setUserCC(ctx, user)

			Convey("Then the result is err", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestGetUserCC(t *testing.T) {
	ctx := context.Background()
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	cc, err := redis.Dial("tcp", s.Addr())

	Convey("Set a user to redis", t, func() {
		adao := &dao{redis: cc}
		user := m.GetUser()
		key := getRedisKey(user.Uid)
		cc.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...)

		Convey("When get this user from redis", func() {
			got, err := adao.getUserCC(ctx, user.Uid)

			Convey("Then the the result is succ", func() {
				So(err, ShouldBeNil)
				So(reflect.DeepEqual(got, user), ShouldBeTrue)
			})
		})

		Convey("When get other user from redis", func() {
			userx := m.GetUser()
			got, err := adao.getUserCC(ctx, userx.Uid)

			Convey("Then the the result is {}", func() {
				So(err, ShouldBeNil)
				So(reflect.DeepEqual(got, &m.User{}), ShouldBeTrue)
			})
		})

		Convey("When close connect, get this user from redis", func() {
			cc.Close()
			_, err := adao.getUserCC(ctx, user.Uid)

			Convey("Then the the result is err", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

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
			if got := getRedisKey(tt.args.uid); got != tt.want {
				t.Errorf("GetRedisKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
