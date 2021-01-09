package misc

import (
	"strconv"
)

// for cache
var expire int64 = 10

//
func GetRedisExpire() int64 {
	return expire
}

//
func SetRedisExpire(time int64) {
	expire = time
}

//
func GetRedisKey(uid int64) string {
	return "uid#" + strconv.FormatInt(uid, 10)
}