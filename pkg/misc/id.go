package misc

import (
	"github.com/sony/sonyflake"
)

func init() {
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
