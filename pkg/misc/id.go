package misc

import (
	"math/rand"
	"time"
)

//
func init() {
	InitGenerator()
	SetRequestIdMax(0x0FFF_FFFF_FFFF_FFFF)
	SetUidMax(0x0FFF_FFFF_FFFF_FFFF)
}

//
func InitGenerator() {
	rand.Seed(time.Now().UnixNano())
}

///////////// request id ///////////////

//
var requestIdMax int64 = 0x0FFF_FFFF_FFFF_FFFF

//
func SetRequestIdMax(max int64) {
	requestIdMax = max
}

//
func GetRequestId() int64 {
	return rand.Int63n(requestIdMax)
}

//////////// user id //////////////
//
var uidmax int64 = 0x0FFF_FFFF_FFFF_FFFF

//
func getUid() int64 {
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
