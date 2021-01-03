package dao

import (
	"reflect"
	"testing"
	"time"

	m "github.com/fuwensun/goms/eTest/internal/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	// New dao
	dao, clean, err := new(getCfgPath())
	if err != nil {
		panic(err)
	}

	// 禁止 log
	level := m.GetLogLevel()

	m.SetLogLevel("")

	testUserCCCRUD(t, dao)
	testUserDBCRUD(t, dao)
	testUserCRUD(t, dao)
	testUserCRUDCacheAside(t, dao)

	// 重置 log
	m.SetLogLevel(level)

	// 清理
	clean()
}

func testUserCCCRUD(t *testing.T, dao *dao) {

	user := m.GetUser()
	Convey("Test dao crud user cc", t, func() {

		Convey("Given a user data for create", func() {
			Convey("When set this user to cache", func() {
				err := dao.setUserCC(ctxb, user)
				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in cc for read", func() {
			Convey("When check this user from cache", func() {
				exist, err := dao.existUserCC(ctxb, user.Uid)
				Convey("Then the result is exist", func() {
					So(err, ShouldBeNil)
					So(exist, ShouldBeTrue)
				})
			})
		})

		Convey("Given a user data in db for update", func() {
			Convey("When get this user from cache", func() {
				got, err := dao.getUserCC(ctxb, user.Uid)
				Convey("Then the result is succ", func() {
					So(reflect.DeepEqual(got, user), ShouldBeTrue)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for delete", func() {
			Convey("When delete this user from cache", func() {
				err := dao.delUserCC(ctxb, user.Uid)
				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Set this user data to redis", func() {
			ex := int64(10)
			inEx := time.Duration(ex/2) * time.Second
			outEx := time.Duration(ex+2) * time.Second
			m.SetExpire(ex)
			user := m.GetUser()
			dao.setUserCC(ctxb, user)
			Convey("When within expiration time", func() {
				time.Sleep(inEx)
				exist, err := dao.existUserCC(ctxb, user.Uid)
				Convey("Then the result is exist", func() {
					So(err, ShouldBeNil)
					So(exist, ShouldBeTrue)
				})
			})

			Convey("When out of expiration time", func() {
				time.Sleep(outEx)
				exist, err := dao.existUserCC(ctxb, user.Uid)
				Convey("Then the result is not exist", func() {
					So(err, ShouldBeNil)
					So(exist, ShouldBeFalse)
				})
			})
		})
	})
}

func testUserDBCRUD(t *testing.T, dao *dao) {

	user := m.GetUser()
	Convey("Test dao crud user db", t, func() {

		Convey("Given a user data for create", func() {
			Convey("When write this user to db", func() {
				err := dao.createUserDB(ctxb, user)
				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for read", func() {
			Convey("When read this user from db", func() {
				got, err := dao.readUserDB(ctxb, user.Uid)
				Convey("Then the result is succ", func() {
					So(reflect.DeepEqual(got, user), ShouldBeTrue)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for update", func() {
			Convey("When update this user to db", func() {
				user.Name = "bar"
				err := dao.updateUserDB(ctxb, user)
				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)
				})
			})
		})
		Convey("Given a user data in db for delete", func() {
			Convey("When delete this user from db", func() {
				err := dao.deleteUserDB(ctxb, user.Uid)
				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func testUserCRUD(t *testing.T, dao *dao) {

	user := m.GetUser()
	Convey("Test dao crud user", t, func() {

		Convey("Given a user data for create", func() {
			Convey("When write this user to dao", func() {
				err := dao.CreateUser(ctxb, user)
				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for read", func() {
			Convey("When read this user from dao", func() {
				got, err := dao.ReadUser(ctxb, user.Uid)
				Convey("Then the result is succ", func() {
					So(reflect.DeepEqual(got, user), ShouldBeTrue)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for update", func() {
			Convey("When update this user to dao", func() {
				user.Name = "bar"
				err := dao.UpdateUser(ctxb, user)
				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for delete", func() {
			Convey("When delete this user from dao", func() {
				err := dao.DeleteUser(ctxb, user.Uid)
				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func testUserCRUDCacheAside(t *testing.T, dao *dao) {

	user := m.GetUser()
	Convey("Test dao read/write user CacheAside", t, func() {

		Convey("Given a user data for create", func() {
			Convey("When write this user to dao", func() {
				err := dao.CreateUser(ctxb, user)
				exist, errcc := dao.existUserCC(ctxb, user.Uid)
				Convey("Then the result is write succ and no cache user data", func() {
					So(err, ShouldBeNil)
					So(errcc, ShouldBeNil)
					So(exist, ShouldBeFalse)
				})
			})
		})

		Convey("Given a user data in db for read", func() {
			Convey("When read this user from dao", func() {
				got, err := dao.ReadUser(ctxb, user.Uid)
				exist, errcc := dao.existUserCC(ctxb, user.Uid)
				Convey("Then the result is read succ and cache user data", func() {
					So(err, ShouldBeNil)
					So(reflect.DeepEqual(got, user), ShouldBeTrue)
					So(errcc, ShouldBeNil)
					So(exist, ShouldBeTrue)
				})
			})
		})

		Convey("Given a user data in db for update", func() {
			Convey("When update this user to dao", func() {
				user.Name = "bar"
				err := dao.UpdateUser(ctxb, user)
				exist, errcc := dao.existUserCC(ctxb, user.Uid)
				Convey("Then the result is succ and delete cached user data", func() {
					So(err, ShouldBeNil)
					So(errcc, ShouldBeNil)
					So(exist, ShouldBeFalse)
				})
			})
		})

		Convey("Given a user data in db for delete", func() {
			dao.CreateUser(ctxb, user)
			dao.ReadUser(ctxb, user.Uid)
			Convey("When delete this user from dao", func() {
				err := dao.DeleteUser(ctxb, user.Uid)
				exist, errcc := dao.existUserCC(ctxb, user.Uid)
				Convey("Then the result is succ and delete cached user data", func() {
					So(err, ShouldBeNil)
					So(errcc, ShouldBeNil)
					So(exist, ShouldBeFalse)
				})
			})
		})
	})
}
