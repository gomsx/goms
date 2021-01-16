package misc

import (
	"github.com/sony/sonyflake"
)

func init() {
	initRidSF()
	initUidSF()
}

//user id
var uidsf *sonyflake.Sonyflake

//
func initUidSF() {
	var st sonyflake.Settings //TODO
	uidsf = sonyflake.NewSonyflake(st)
	if uidsf == nil {
		panic("sonyflake not created")
	}
}

//
func GenUid() int64 {
	id, err := uidsf.NextID()
	if err != nil {
		panic("uid not generate")
	}
	return int64(id)
}

//
func getUid() int64 {
	return GenUid()
}

//request id
var ridsf *sonyflake.Sonyflake

//
func initRidSF() {
	var st sonyflake.Settings //TODO
	ridsf = sonyflake.NewSonyflake(st)
	if ridsf == nil {
		panic("sonyflake not created")
	}
}

//
func GenRequestId() int64 {
	id, err := ridsf.NextID()
	if err != nil {
		panic("rid not generate")
	}
	return int64(id)
}

//
func GetRequestId() int64 {
	return GenRequestId()
}
