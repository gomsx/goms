package ecode

import (
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {

	en := 123
	es := strconv.FormatInt(int64(en), 10)
	emsg := es
	ec := New(en)

	if ec.Code() != en {
		t.Errorf("got = %v, want %v", ec.Code(), en)
	}

	if ec.String() != es {
		t.Errorf("got = %v, want %v", ec.String(), es)
	}

	if ec.Msg() != emsg {
		t.Errorf("got = %v, want %v", ec.Msg(), emsg)
	}

	if _, ok := __codes[en]; ok != true {
		t.Errorf("got = %v, want %v", ok, true)
	}

	en2 := 456
	New(en2)
	if _, ok := __codes[en2]; ok != true {
		t.Errorf("got = %v, want %v", ok, true)
	}

	defer func() {
		if p := recover(); p == nil {
			t.Errorf("got nil, want panic")
		}
	}()
	New(en2)

}

