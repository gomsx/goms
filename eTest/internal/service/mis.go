package service

import (
	"math/rand"
	"time"
)

func initUidGenerator() {
	rand.Seed(time.Now().UnixNano())
}
func getUid() int64 {
	return rand.Int63n(0x0FFF_FFFF_FFFF_FFFF)
}
