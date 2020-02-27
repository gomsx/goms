package model

import "golang.org/x/exp/errors"

var ErrNotFound = errors.New("not found!")
var ErrUidExsit = errors.New("uid exist!")
