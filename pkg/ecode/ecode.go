package ecode

import (
	"fmt"
	"strconv"
)

var (
	__codes = map[int]struct{}{}
	__msgs  = map[int]string{}
)

type Ecode interface {
	String() string
	Code() int
	Msg() string
}

type Code int64

func (e Code) String() string {
	return strconv.FormatInt(int64(e), 10)
}
func (e Code) Code() int { return int(e) }

func (e Code) Msg() string {
	if msg, ok := __msgs[e.Code()]; ok {
		return msg
	}
	return e.String()
}

func New(e int) Code {
	if _, ok := __codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	__codes[e] = struct{}{}
	return Code(e)
}
