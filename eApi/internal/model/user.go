package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/exp/errors"
)

func init() {
	InitUidGenerator()
	SetUidMax(0x0FFF_FFFF_FFFF_FFFF)
}

type User struct {
	Uid  int64  `redis:"uid"`
	Name string `redis:"name"`
	Sex  int64  `redis:"sex"`
}

//
func GetRedisKey(uid int64) string {
	return "uid#" + strconv.FormatInt(uid, 10)
}

//
var uidmax int64 = 0x0FFF_FFFF_FFFF_FFFF

func InitUidGenerator() {
	rand.Seed(time.Now().UnixNano())
}
func GetUid() int64 {
	return rand.Int63n(uidmax)
}
func SetUidMax(max int64) {
	uidmax = max
}
func GetUidMax() int64 {
	return uidmax
}

//
func CheckUid(uid int64) bool {
	var min int64 = 0
	var max int64 = uidmax
	if uid >= min && uid <= max {
		return true
	}
	return false
}
func CheckName(name string) bool {
	var min int = 1
	var max int = 18
	if len(name) >= min && len(name) <= max {
		return true
	}
	return false
}
func CheckSex(sex int64) bool {
	var min int64 = 0
	var max int64 = 1
	if sex >= min && sex <= max {
		return true
	}
	return false
}
func CheckUidS(uidstr string) (int64, bool) {
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if err != nil {
		return -1, false
	}
	return uid, CheckUid(uid)
}
func CheckSexS(sexstr string) (int64, bool) {
	sex, err := strconv.ParseInt(sexstr, 10, 64)
	if err != nil {
		return -1, false
	}
	return sex, CheckSex(sex)
}

var ErrArgError = errors.New("arg error")

var ErrUidError = fmt.Errorf("uid %w, ECODE-%d", ErrArgError, 10001)
var ErrNameError = fmt.Errorf("name %w, ECODE-%d", ErrArgError, 10002)
var ErrSexError = fmt.Errorf("sex %w, ECODE-%d", ErrArgError, 10003)

