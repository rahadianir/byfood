package xerrors

import "fmt"

var (
	ErrDataNotFound = fmt.Errorf("data not found")
	ErrInvalidID    = fmt.Errorf("invalid id")
)
