package dao

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	m "github.com/gomsx/goms/eApi/internal/model"

	sm "github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
)

var dbdao *dao
var dbmock sm.Sqlmock
var dbconn *sql.DB

//
func tearupDb() {
	var err error
	dbconn, dbmock, err = sm.New()
	if err != nil {
		panic(err)
	}
	dbdao = &dao{db: dbconn}
}

//
func teardownDb() {
	dbconn.Close()
}

func TestCreateUserDB(t *testing.T) {
	user := m.GetUser()
	createUser := "INSERT INTO user_table"

	Convey("Create user when succeed to operate database", t, func() {
		dbmock.ExpectExec(createUser).
			WithArgs(user.Uid, user.Name, user.Sex).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(nil)

		err := dbdao.createUserDB(ctxb, user)
		So(err, ShouldBeNil)
	})

	Convey("Create user when fail to operate database", t, func() {
		dbmock.ExpectExec(createUser).
			WithArgs(user.Uid, user.Name, user.Sex).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(errx)

		err := dbdao.createUserDB(ctxb, user)
		So(errors.Is(err, errx), ShouldBeTrue)
	})
}

func TestReadUserDB(t *testing.T) {
	user := m.GetUser()

	Convey("Read user when succeed to operate database", t, func() {
		rows := sm.NewRows([]string{"uid", "name", "sex"}).
			AddRow(user.Uid, user.Name, user.Sex)

		dbmock.ExpectQuery(_readUser).
			WithArgs(user.Uid).
			WillReturnRows(rows).
			WillReturnError(nil)

		got, err := dbdao.readUserDB(ctxb, user.Uid)
		So(err, ShouldBeNil)
		So(reflect.DeepEqual(got, user), ShouldBeTrue)
	})

	Convey("Read user when fail to operate database", t, func() {
		dbmock.ExpectQuery(_readUser).
			WithArgs(user.Uid).
			WillReturnRows(nil).
			WillReturnError(errx)

		_, err := dbdao.readUserDB(ctxb, user.Uid)
		So(errors.Is(err, errx), ShouldBeTrue)
	})
}

func TestUpdateUserDB(t *testing.T) {
	user := m.GetUser()
	updateUser := "UPDATE user_table"

	Convey("Update user when succeed to operate database", t, func() {
		dbmock.ExpectExec(updateUser).
			WithArgs(user.Name, user.Sex, user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(nil)

		err := dbdao.updateUserDB(ctxb, user)
		So(err, ShouldBeNil)
	})

	Convey("Update user when fail to operate database", t, func() {
		dbmock.ExpectExec(updateUser).
			WithArgs(user.Name, user.Sex, user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(errx)

		err := dbdao.updateUserDB(ctxb, user)
		So(errors.Is(err, errx), ShouldBeTrue)
	})
}

func TestDeleteUserDB(t *testing.T) {
	user := m.GetUser()

	Convey("Delete user when succeed to operate database", t, func() {
		dbmock.ExpectExec(_deleteUser).
			WithArgs(user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(nil)

		err := dbdao.deleteUserDB(ctxb, user.Uid)
		So(err, ShouldBeNil)
	})

	Convey("Delete user when fail to operate database", t, func() {
		dbmock.ExpectExec(_deleteUser).
			WithArgs(user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(errx)

		err := dbdao.deleteUserDB(ctxb, user.Uid)
		So(errors.Is(err, errx), ShouldBeTrue)
	})
}
