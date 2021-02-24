package err

import (
	"errors"
	"fmt"
)

// data
var ErrNotFound = errors.New("not found")
var ErrNotFoundData = fmt.Errorf("data %w", ErrNotFound)

var ErrFailedCreate = errors.New("failed to create")
var ErrFailedCreateData = fmt.Errorf("data %w", ErrFailedCreate)

var ErrInternalError = errors.New("internal error")

// user
var ErrArgError = errors.New("arg error")

var ErrUidError = fmt.Errorf("uid %w", ErrArgError)
var ErrNameError = fmt.Errorf("name %w", ErrArgError)
var ErrSexError = fmt.Errorf("sex %w", ErrArgError)

var UserErrMap = map[string]error{
	"Uid":  ErrUidError,
	"Name": ErrNameError,
	"Sex":  ErrSexError,
}

var UserEcodeMap = map[string]int64{
	"Uid":  10001,
	"Name": 10002,
	"Sex":  10003,
}

const (
	StatusOK                  = 200 //
	StatusBadRequest          = 400 //
	StatusInternalServerError = 500 //
)
