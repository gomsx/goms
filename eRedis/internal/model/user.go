package model

import "strconv"

type User struct {
	Uid  int64  `redis:"uid"`
	Name string `redis:"name"`
	Sex  int64  `redis:"sex"`
}

func GetRedisKey(uid int64) string {
	return "uid#" + strconv.FormatInt(uid, 10)
}

