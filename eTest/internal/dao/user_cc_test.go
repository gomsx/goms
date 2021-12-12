package dao

import (
	"reflect"
	"testing"

	m "github.com/gomsx/goms/eTest/internal/model"

	rm "github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
)

var ccdao *dao
var ccmock *rm.Miniredis
var ccconn redis.Conn

//
func tearupCache() {
	var err error
	ccmock, err = rm.Run()
	if err != nil {
		panic(err)
	}
	ccconn, err = redis.Dial("tcp", ccmock.Addr())
	ccdao = &dao{redis: ccconn}
}

//
func teardownCache() {
	ccmock.Close()
}
func TestExistUserCC(t *testing.T) {
	user := m.GetUser()

	Convey("Given a user in redis", t, func() {
		key := getRedisKey(user.Uid)
		ccconn.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...)

		Convey("When check this user from redis", func() {
			exist, err := ccdao.existUserCC(ctxb, user.Uid)

			Convey("Then the result should exist", func() {
				So(err, ShouldBeNil)
				So(exist, ShouldBeTrue)
			})
		})

		Convey("When check other user from redis", func() {
			userx := m.GetUser()
			exist, err := ccdao.existUserCC(ctxb, userx.Uid)

			Convey("Then the result should not exist", func() {
				So(err, ShouldBeNil)
				So(exist, ShouldBeFalse)
			})
		})
	})
}

func TestSetUserCC(t *testing.T) {
	user := m.GetUser()

	Convey("Given a user data", t, func() {

		Convey("When set this user to redis", func() {
			err := ccdao.setUserCC(ctxb, user)

			Convey("Then the result should succeed", func() {
				So(err, ShouldBeNil)

				Convey("When set same user to redis", func() {
					err := ccdao.setUserCC(ctxb, user)

					Convey("Then the result should succeed", func() {
						So(err, ShouldBeNil)
					})
				})
			})
		})

		Convey("When set other user to redis", func() {
			userx := m.GetUser()
			err := ccdao.setUserCC(ctxb, userx)

			Convey("Then the result should succeed", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestGetUserCC(t *testing.T) {
	user := m.GetUser()

	Convey("Given a user in redis", t, func() {
		key := getRedisKey(user.Uid)
		ccconn.Do("HMSET", redis.Args{}.Add(key).AddFlat(user)...)

		Convey("When get this user from redis", func() {
			got, err := ccdao.getUserCC(ctxb, user.Uid)

			Convey("Then the result should succeed", func() {
				So(err, ShouldBeNil)
				So(reflect.DeepEqual(got, user), ShouldBeTrue)
			})
		})

		Convey("When get other user from redis", func() {
			userx := m.GetUser()
			got, err := ccdao.getUserCC(ctxb, userx.Uid)

			Convey("Then the result should be {}", func() {
				So(err, ShouldBeNil)
				So(reflect.DeepEqual(got, &m.User{}), ShouldBeTrue)
			})
		})
	})
}
