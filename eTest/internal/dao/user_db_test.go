package dao

import (
	"database/sql"
	"reflect"
	"testing"

	m "github.com/aivuca/goms/eTest/internal/model"

	sm "github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/exp/errors"
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

	Convey("Test CreateUserDB succ", t, func() {
		dbmock.ExpectExec(createUser).
			WithArgs(user.Uid, user.Name, user.Sex).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(nil)

		err := dbdao.createUserDB(ctxb, user)
		So(err, ShouldBeNil)
	})

	Convey("Test CreateUserDB fail", t, func() {
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

	Convey("Test ReadUserDB succ", t, func() {
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

	Convey("Test ReadUserDB fail", t, func() {
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

	Convey("Test UpdateUserDB succ", t, func() {
		dbmock.ExpectExec(updateUser).
			WithArgs(user.Name, user.Sex, user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(nil)

		err := dbdao.updateUserDB(ctxb, user)
		So(err, ShouldBeNil)
	})

	Convey("Test UpdateUserDB fail", t, func() {
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

	Convey("Test DeleteUserDB succ", t, func() {
		dbmock.ExpectExec(_deleteUser).
			WithArgs(user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(nil)

		err := dbdao.deleteUserDB(ctxb, user.Uid)
		So(err, ShouldBeNil)
	})

	Convey("Test DeleteUserDB fail", t, func() {
		dbmock.ExpectExec(_deleteUser).
			WithArgs(user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(errx)

		err := dbdao.deleteUserDB(ctxb, user.Uid)
		So(errors.Is(err, errx), ShouldBeTrue)
	})
}
