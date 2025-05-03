package sharederr

import "errors"
var (
ErrNotFound = errors.New("not found")
ErrDuplicate = errors.New("duplicate")
ErrMismatchedPass = errors.New("mismatched password")
)
