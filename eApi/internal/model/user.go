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
	Uid  int64  `redis:"uid" validate:"required,gte=0"`
	Name string `redis:"name" validate:"required,min=1,max=18"`
	Sex  int64  `redis:"sex" validate:"required,gte=1,lte=2"`
}

var UserErrMap = map[string]error{
	"User.Uid":  ErrUidError,
	"User.Name": ErrNameError,
	"User.Sex":  ErrSexError,
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

var ErrArgError = errors.New("arg error")

var ErrUidError = fmt.Errorf("uid %w, ECODE-%d", ErrArgError, 10001)
var ErrNameError = fmt.Errorf("name %w, ECODE-%d", ErrArgError, 10002)
var ErrSexError = fmt.Errorf("sex %w, ECODE-%d", ErrArgError, 10003)
