package dao

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	m "github.com/aivuca/goms/eTest/internal/model"

	sm "github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/exp/errors"
)

var adao *dao
var adb *sql.DB
var asm sm.Sqlmock
var xerr = fmt.Errorf("test error")
var xctx = context.Background()

//
func tearupSqlmock() {
	var err error
	adb, asm, err = sm.New()
	if err != nil {
		panic("failed to tear up sqlmock")
	}
	adao = &dao{db: adb}
}

//
func teardownSqlmock() {
	adb.Close()
}

func TestCreateUserDB(t *testing.T) {
	user := m.GetUser()
	createUser := "INSERT INTO user_table"

	Convey("Test CreateUserDB succ", t, func() {
		asm.ExpectExec(createUser).
			WithArgs(user.Uid, user.Name, user.Sex).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(nil)

		err := adao.createUserDB(xctx, user)
		So(err, ShouldBeNil)
	})

	Convey("Test CreateUserDB fail", t, func() {
		asm.ExpectExec(createUser).
			WithArgs(user.Uid, user.Name, user.Sex).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(xerr)

		err := adao.createUserDB(xctx, user)
		So(errors.Is(err, xerr), ShouldBeTrue)
	})
}

func TestReadUserDB(t *testing.T) {
	user := m.GetUser()

	Convey("Test ReadUserDB succ", t, func() {
		rows := sm.NewRows([]string{"uid", "name", "sex"}).
			AddRow(user.Uid, user.Name, user.Sex)

		asm.ExpectQuery(_readUser).
			WithArgs(user.Uid).
			WillReturnRows(rows).
			WillReturnError(nil)

		got, err := adao.readUserDB(xctx, user.Uid)
		So(err, ShouldBeNil)
		So(reflect.DeepEqual(got, user), ShouldBeTrue)
	})

	Convey("Test ReadUserDB fail", t, func() {
		asm.ExpectQuery(_readUser).
			WithArgs(user.Uid).
			WillReturnRows(nil).
			WillReturnError(xerr)

		_, err := adao.readUserDB(xctx, user.Uid)
		So(errors.Is(err, xerr), ShouldBeTrue)
	})
}

func TestUpdateUserDB(t *testing.T) {
	user := m.GetUser()
	updateUser := "UPDATE user_table"

	Convey("Test UpdateUserDB succ", t, func() {
		asm.ExpectExec(updateUser).
			WithArgs(user.Name, user.Sex, user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(nil)

		err := adao.updateUserDB(xctx, user)
		So(err, ShouldBeNil)
	})

	Convey("Test UpdateUserDB fail", t, func() {
		asm.ExpectExec(updateUser).
			WithArgs(user.Name, user.Sex, user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(xerr)

		err := adao.updateUserDB(xctx, user)
		So(errors.Is(err, xerr), ShouldBeTrue)
	})
}

func TestDeleteUserDB(t *testing.T) {
	user := m.GetUser()

	Convey("Test DeleteUserDB succ", t, func() {
		asm.ExpectExec(_deleteUser).
			WithArgs(user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(nil)

		err := adao.deleteUserDB(xctx, user.Uid)
		So(err, ShouldBeNil)
	})

	Convey("Test DeleteUserDB fail", t, func() {
		asm.ExpectExec(_deleteUser).
			WithArgs(user.Uid).
			WillReturnResult(sm.NewResult(1, 1)).
			WillReturnError(xerr)

		err := adao.deleteUserDB(xctx, user.Uid)
		So(errors.Is(err, xerr), ShouldBeTrue)
	})
}
