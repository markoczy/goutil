package clierror

import (
	"fmt"
)

const strArgsMismatch = "Argument count mismatch"

// ErrorArgsCountMismatch ...
var ErrorArgsCountMismatch = fmt.Errorf(strArgsMismatch)

// IsArgsCountMismatch ...
func IsArgsCountMismatch(err error) bool {
	return err.Error() == strArgsMismatch
}
