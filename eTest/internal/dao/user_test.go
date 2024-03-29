package dao

import (
	"reflect"
	"testing"

	m "github.com/gomsx/goms/eTest/internal/model"
	"github.com/spf13/viper"

	. "github.com/smartystreets/goconvey/convey"
)

var udao *dao
var clean func()

//
func tearupDao() {
	var err error

	viper.Reset()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(getCfgPath())
	viper.ReadInConfig()

	udao, clean, err = new()
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
				Convey("Then the result should exist", func() {
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
				Convey("Then the result should write succeeded and no cache user data", func() {
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
				Convey("Then the result should read succeeded and cache user data", func() {
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
