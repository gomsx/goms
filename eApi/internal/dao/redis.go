package dao

import (
	"strconv"
)

// for cache
var expire int64 = 10

//
func getRedisExpire() int64 {
	return expire
}

//
func setRedisExpire(time int64) {
	expire = time
}

//
func getRedisKey(uid int64) string {
	return "uid#" + strconv.FormatInt(uid, 16)
}
