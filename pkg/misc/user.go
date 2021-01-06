package misc

import (
	"strconv"
)

//
func GetUid() int64 {
	return getUid()
}

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
func StrInt(intx int64) string {
	return strconv.FormatInt(intx, 10)
}
