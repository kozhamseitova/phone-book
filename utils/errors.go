package utils

import "errors"

var (
	ErrInternalError = errors.New("internal server error")
	ErrInvalidParam = errors.New("invalid param")
)