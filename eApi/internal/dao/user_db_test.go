package dao

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	m "github.com/aivuca/goms/eApi/internal/model"

	smk "github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/exp/errors"
)

var daox *dao
var dbx *sql.DB
var smx smk.Sqlmock
var errx = fmt.Errorf("test error")

//
func tearupSqlmock() {
	var err error
	dbx, smx, err = smk.New()
	if err != nil {
	}
	daox = &dao{db: dbx}
}

//
func teardownSqlmock() {
	dbx.Close()
}

func Test_CreateUserDB(t *testing.T) {
	user := m.GetUser()
	createUser := "INSERT INTO user_table"

	Convey("Test CreateUserDB succ", t, func() {
		smx.ExpectExec(createUser).
			WithArgs(user.Uid, user.Name, user.Sex).
			WillReturnResult(smk.NewResult(1, 1)).
			WillReturnError(nil)

		err := daox.createUserDB(ctx, user)
		So(err, ShouldBeNil)
	})

	Convey("Test CreateUserDB fail", t, func() {
		smx.ExpectExec(createUser).
			WithArgs(user.Uid, user.Name, user.Sex).
			WillReturnResult(smk.NewResult(1, 1)).
			WillReturnError(errx)

		err := daox.createUserDB(ctx, user)
		So(errors.Is(err, errx), ShouldBeTrue)
	})
}

func Test_ReadUserDB(t *testing.T) {
	user := m.GetUser()

	Convey("Test ReadUserDB succ", t, func() {
		rows := smk.NewRows([]string{"uid", "name", "sex"}).
			AddRow(user.Uid, user.Name, user.Sex)

		smx.ExpectQuery(_readUser).
			WithArgs(user.Uid).
			WillReturnRows(rows).
			WillReturnError(nil)

		got, err := daox.readUserDB(ctx, user.Uid)
		So(err, ShouldBeNil)
		So(reflect.DeepEqual(got, user), ShouldBeTrue)
	})

	Convey("Test ReadUserDB fail", t, func() {
		smx.ExpectQuery(_readUser).
			WithArgs(user.Uid).
			WillReturnRows(nil).
			WillReturnError(errx)

		_, err := daox.readUserDB(ctx, user.Uid)
		So(errors.Is(err, errx), ShouldBeTrue)
	})
}

func Test_UpdateUserDB(t *testing.T) {
	user := m.GetUser()
	updateUser := "UPDATE user_table"

	Convey("Test UpdateUserDB succ", t, func() {
		smx.ExpectExec(updateUser).
			WithArgs(user.Name, user.Sex, user.Uid).
			WillReturnResult(smk.NewResult(1, 1)).
			WillReturnError(nil)

		err := daox.updateUserDB(ctx, user)
		So(err, ShouldBeNil)
	})

	Convey("Test UpdateUserDB fail", t, func() {
		smx.ExpectExec(updateUser).
			WithArgs(user.Name, user.Sex, user.Uid).
			WillReturnResult(smk.NewResult(1, 1)).
			WillReturnError(errx)

		err := daox.updateUserDB(ctx, user)
		So(errors.Is(err, errx), ShouldBeTrue)
	})
}

func Test_DeleteUserDB(t *testing.T) {
	user := m.GetUser()

	Convey("Test DeleteUserDB succ", t, func() {
		smx.ExpectExec(_deleteUser).
			WithArgs(user.Uid).
			WillReturnResult(smk.NewResult(1, 1)).
			WillReturnError(nil)

		err := daox.deleteUserDB(ctx, user.Uid)
		So(err, ShouldBeNil)
	})

	Convey("Test DeleteUserDB fail", t, func() {
		smx.ExpectExec(_deleteUser).
			WithArgs(user.Uid).
			WillReturnResult(smk.NewResult(1, 1)).
			WillReturnError(errx)

		err := daox.deleteUserDB(ctx, user.Uid)
		So(errors.Is(err, errx), ShouldBeTrue)
	})
}
