package service

import (
	"math/rand"
	"time"
)

var uidmax int64 = 0x0FFF_FFFF_FFFF_FFFF

func initUidGenerator() {
	rand.Seed(time.Now().UnixNano())
}
func getUid() int64 {
	return rand.Int63n(uidmax)
}
func setUidMax(max int64) {
	uidmax = max
}
func getUidMax() int64 {
	return uidmax
}
