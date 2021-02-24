package misc

import (
	"strconv"
)

//
func GetUid() int64 {
	return GenUid()
}

//
func GetUidBad() int64 {
	return -1 * GetUid()
}

var cnt int64 = 0

//
func GetName() string {
	cnt++
	return "name_test" + "_" + strconv.FormatInt(cnt, 10)
}

//
func GetNameBad() string {
	return GetName() + "_bad" + "_&%$!@*?"
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
func StrInt(intx int64) string {
	return strconv.FormatInt(intx, 10)
}
