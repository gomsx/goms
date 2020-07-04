package model

import (
	"fmt"

	"golang.org/x/exp/errors"
)

var ErrNotFound = errors.New("not found")
var ErrNotFoundData = fmt.Errorf("data:%w", ErrNotFound)

var ErrFailedCreate = errors.New("failed to create")
var ErrFailedCreateData = fmt.Errorf("data:%w", ErrFailedCreate)

var ErrInternalError = errors.New("internal error")

var CfgPath = ""

