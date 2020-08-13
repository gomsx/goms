package dao

import (
	"fmt"
	"reflect"
	"testing"

	m "github.com/aivuca/goms/eApi/internal/model"
	lg "github.com/aivuca/goms/eApi/internal/pkg/log"

	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	if isCiEnvDocker() {
		cpstub := gostub.Stub(&cfgpath, "testdata/teardocker/configs")
		defer cpstub.Reset()
		fmt.Printf("stub config path to: %v", cfgpath)
	}

	// New dao
	dao, clean, err := new(cfgpath)
	if err != nil {
		panic(err)
	}

	level := lg.GetLevel()
	lg.SetLevel("")

	Convey("Test dao crud user", t, func() {

		Convey("Given a user data", func() {
			user := m.GetUser()

			Convey("When write this user to dao", func() {
				err := dao.CreateUser(ctx, user)

				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)

					Convey("When read this user from dao", func() {
						got, err := dao.ReadUser(ctx, user.Uid)

						Convey("Then the result is succ", func() {
							So(reflect.DeepEqual(got, user), ShouldBeTrue)
							So(err, ShouldBeNil)

							Convey("When update this user to dao", func() {
								user.Name = "bar"
								err = dao.UpdateUser(ctx, user)

								Convey("Then the result is succ", func() {
									So(err, ShouldBeNil)

									Convey("When delete this user from dao", func() {
										err = dao.DeleteUser(ctx, user.Uid)

										Convey("Then the result is succ", func() {
											So(err, ShouldBeNil)
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})

	Convey("Test dao crud user db", t, func() {

		Convey("Given a user data", func() {
			user := m.GetUser()

			Convey("When write this user to db", func() {
				err := dao.createUserDB(ctx, user)

				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)

					Convey("When read this user from db", func() {
						got, err := dao.readUserDB(ctx, user.Uid)

						Convey("Then the result is succ", func() {
							So(reflect.DeepEqual(got, user), ShouldBeTrue)
							So(err, ShouldBeNil)

							Convey("When update this user to db", func() {
								user.Name = "bar"
								err = dao.updateUserDB(ctx, user)

								Convey("Then the result is succ", func() {
									So(err, ShouldBeNil)

									Convey("When delete this user from db", func() {
										err = dao.deleteUserDB(ctx, user.Uid)

										Convey("Then the result is succ", func() {
											So(err, ShouldBeNil)
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})

	Convey("Test dao crud user cc", t, func() {

		Convey("Given a user data", func() {
			user := m.GetUser()

			Convey("When set this user to cache", func() {
				err := dao.setUserCC(ctx, user)

				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)

					Convey("When check this user from cache", func() {
						exist, err := dao.existUserCC(ctx, user.Uid)

						Convey("Then the result is exist", func() {
							So(err, ShouldBeNil)
							So(exist, ShouldBeTrue)

							Convey("When get this user from cache", func() {
								got, err := dao.getUserCC(ctx, user.Uid)

								Convey("Then the result is succ", func() {
									So(reflect.DeepEqual(got, user), ShouldBeTrue)
									So(err, ShouldBeNil)

									Convey("When delete this user from cache", func() {
										err = dao.delUserCC(ctx, user.Uid)

										Convey("Then the result is succ", func() {
											So(err, ShouldBeNil)

											Convey("When check this user from cache", func() {
												exist, err = dao.existUserCC(ctx, user.Uid)

												Convey("Then the result is not exist", func() {
													So(err, ShouldBeNil)
													So(exist, ShouldBeFalse)
												})
											})
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})

	Convey("Test dao read/write user Cache-aside", t, func() {

		Convey("Given a user data", func() {
			user := m.GetUser()

			Convey("When write this user to dao", func() {
				err := dao.CreateUser(ctx, user)

				Convey("Then the result is succ", func() {
					So(err, ShouldBeNil)

					Convey("When check this user from cache", func() {
						exist, err := dao.existUserCC(ctx, user.Uid)

						Convey("Then the result is not exist", func() {
							So(err, ShouldBeNil)
							So(exist, ShouldBeFalse)

							Convey("When read this user from dao", func() {
								got, err := dao.ReadUser(ctx, user.Uid)

								Convey("Then the result is succ", func() {
									So(reflect.DeepEqual(got, user), ShouldBeTrue)
									So(err, ShouldBeNil)

									Convey("When check this user from cache", func() {
										exist, err = dao.existUserCC(ctx, user.Uid)

										Convey("Then the result is exist", func() {
											So(err, ShouldBeNil)
											So(exist, ShouldBeTrue)

											Convey("When delete this user from dao", func() {
												err = dao.DeleteUser(ctx, user.Uid)

												Convey("Then the result is succ", func() {
													So(err, ShouldBeNil)

													Convey("When check this user from cache", func() {
														exist, err = dao.existUserCC(ctx, user.Uid)

														Convey("Then the result is not exist", func() {
															So(err, ShouldBeNil)
															So(exist, ShouldBeFalse)
														})
													})
												})
											})

											Convey("When update this user to dao", func() {
												user.Name = "bar"
												err = dao.UpdateUser(ctx, user)

												Convey("Then the result is succ", func() {
													So(err, ShouldBeNil)

													Convey("When check this user from cache", func() {
														exist, err = dao.existUserCC(ctx, user.Uid)

														Convey("Then the result is not exist", func() {
															So(err, ShouldBeNil)
															So(exist, ShouldBeFalse)
														})
													})
												})
											})
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})

	lg.SetLevel(level)

	// 清理
	clean()
}