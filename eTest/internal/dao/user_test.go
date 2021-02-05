package dao

import (
	"reflect"
	"testing"
	"time"

	m "github.com/aivuca/goms/eTest/internal/model"
	ms "github.com/aivuca/goms/pkg/misc"

	. "github.com/smartystreets/goconvey/convey"
)

var udao *dao
var clean func()

//
func tearupDao() {
	var err error
	udao, clean, err = new(getCfgPath())
	if err != nil {
		panic(err)
	}
}

//
func teardownDao() {
	clean()
}

func TestUserCCCRUD(t *testing.T) {

	user := m.GetUser()
	Convey("Test dao crud user cc", t, func() {

		Convey("Given a user data for create", func() {
			Convey("When set this user to cache", func() {
				err := udao.setUserCC(ctxb, user)
				Convey("Then the result should succeed", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in cc for read", func() {
			Convey("When check this user from cache", func() {
				exist, err := udao.existUserCC(ctxb, user.Uid)
				Convey("Then the result should existed", func() {
					So(err, ShouldBeNil)
					So(exist, ShouldBeTrue)
				})
			})
		})

		Convey("Given a user data in db for update", func() {
			Convey("When get this user from cache", func() {
				got, err := udao.getUserCC(ctxb, user.Uid)
				Convey("Then the result should succeed", func() {
					So(reflect.DeepEqual(got, user), ShouldBeTrue)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for delete", func() {
			Convey("When delete this user from cache", func() {
				err := udao.delUserCC(ctxb, user.Uid)
				Convey("Then the result should succeed", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Set this user data to redis", func() {
			ex := int64(10)
			inEx := time.Duration(ex/2) * time.Second
			outEx := time.Duration(ex+2) * time.Second
			ms.SetRedisExpire(ex)
			user := m.GetUser()
			udao.setUserCC(ctxb, user)
			Convey("When within expiration time", func() {
				time.Sleep(inEx)
				exist, err := udao.existUserCC(ctxb, user.Uid)
				Convey("Then the result should existed", func() {
					So(err, ShouldBeNil)
					So(exist, ShouldBeTrue)
				})
			})

			Convey("When out of expiration time", func() {
				time.Sleep(outEx)
				exist, err := udao.existUserCC(ctxb, user.Uid)
				Convey("Then the result should not existed", func() {
					So(err, ShouldBeNil)
					So(exist, ShouldBeFalse)
				})
			})
		})
	})
}

func TestUserDBCRUD(t *testing.T) {

	user := m.GetUser()
	Convey("Test dao crud user db", t, func() {

		Convey("Given a user data for create", func() {
			Convey("When write this user to db", func() {
				err := udao.createUserDB(ctxb, user)
				Convey("Then the result should succeed", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for read", func() {
			Convey("When read this user from db", func() {
				got, err := udao.readUserDB(ctxb, user.Uid)
				Convey("Then the result should succeed", func() {
					So(reflect.DeepEqual(got, user), ShouldBeTrue)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for update", func() {
			Convey("When update this user to db", func() {
				user.Name = "bar"
				err := udao.updateUserDB(ctxb, user)
				Convey("Then the result should succeed", func() {
					So(err, ShouldBeNil)
				})
			})
		})
		Convey("Given a user data in db for delete", func() {
			Convey("When delete this user from db", func() {
				err := udao.deleteUserDB(ctxb, user.Uid)
				Convey("Then the result should succeed", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestUserCRUD(t *testing.T) {

	user := m.GetUser()
	Convey("Test dao crud user", t, func() {

		Convey("Given a user data for create", func() {
			Convey("When write this user to dao", func() {
				err := udao.CreateUser(ctxb, user)
				Convey("Then the result should succeed", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for read", func() {
			Convey("When read this user from dao", func() {
				got, err := udao.ReadUser(ctxb, user.Uid)
				Convey("Then the result should succeed", func() {
					So(reflect.DeepEqual(got, user), ShouldBeTrue)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for update", func() {
			Convey("When update this user to dao", func() {
				user.Name = "bar"
				err := udao.UpdateUser(ctxb, user)
				Convey("Then the result should succeed", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("Given a user data in db for delete", func() {
			Convey("When delete this user from dao", func() {
				err := udao.DeleteUser(ctxb, user.Uid)
				Convey("Then the result should succeed", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestUserCRUDCacheAside(t *testing.T) {

	user := m.GetUser()
	Convey("Test dao read/write user CacheAside", t, func() {

		Convey("Given a user data for create", func() {
			Convey("When write this user to dao", func() {
				err := udao.CreateUser(ctxb, user)
				exist, errcc := udao.existUserCC(ctxb, user.Uid)
				Convey("Then the result should be write succ and no cache user data", func() {
					So(err, ShouldBeNil)
					So(errcc, ShouldBeNil)
					So(exist, ShouldBeFalse)
				})
			})
		})

		Convey("Given a user data in db for read", func() {
			Convey("When read this user from dao", func() {
				got, err := udao.ReadUser(ctxb, user.Uid)
				exist, errcc := udao.existUserCC(ctxb, user.Uid)
				Convey("Then the result should be read succ and cache user data", func() {
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
				err := udao.UpdateUser(ctxb, user)
				exist, errcc := udao.existUserCC(ctxb, user.Uid)
				Convey("Then the result should succeed and delete cached user data", func() {
					So(err, ShouldBeNil)
					So(errcc, ShouldBeNil)
					So(exist, ShouldBeFalse)
				})
			})
		})

		Convey("Given a user data in db for delete", func() {
			udao.CreateUser(ctxb, user)
			udao.ReadUser(ctxb, user.Uid)
			Convey("When delete this user from dao", func() {
				err := udao.DeleteUser(ctxb, user.Uid)
				exist, errcc := udao.existUserCC(ctxb, user.Uid)
				Convey("Then the result should succeed and delete cached user data", func() {
					So(err, ShouldBeNil)
					So(errcc, ShouldBeNil)
					So(exist, ShouldBeFalse)
				})
			})
		})
	})
}
