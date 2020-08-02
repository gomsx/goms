package model

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	InitUidGenerator()
	SetUidMax(0x0FFF_FFFF_FFFF_FFFF)
}

//
type User struct {
	Uid  int64  `redis:"uid" validate:"required,gte=0"`
	Name string `redis:"name" validate:"required,min=1,max=18"`
	Sex  int64  `redis:"sex" validate:"required,gte=1,lte=2"`
}

//
var uidmax int64 = 0x0FFF_FFFF_FFFF_FFFF

//
func InitUidGenerator() {
	rand.Seed(time.Now().UnixNano())
}

//
func GetUid() int64 {
	return rand.Int63n(uidmax)
}

//
func SetUidMax(max int64) {
	uidmax = max
}

//
func GetUidMax() int64 {
	return uidmax
}

//
func GetRedisKey(uid int64) string {
	return "uid#" + strconv.FormatInt(uid, 10)
}

// for test
//
func GetUidBad() int64 {
	return -1 * GetUid()
}

//
func GetName() string {
	return "namexxx"
}

//
func GetNameBad() string {
	return GetName() + "&%$!@*?"
}

//
func GetSex() int64 {
	return 1
}

//
func GetSexBad() int64 {
	return GetSex() + 100000
}

//
func StrInt(sex int64) string {
	return strconv.FormatInt(sex, 10)
}
