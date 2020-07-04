package model

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found!")
var ErrNotFoundData = fmt.Errorf("data:%w", ErrNotFound)

var ErrFailedCreate = fmt.Errorf("failed to create!")
var ErrFailedCreateData = fmt.Errorf("data:%w", ErrFailedCreate)

