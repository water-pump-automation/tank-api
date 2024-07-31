package get_group

import (
	"errors"
	"fmt"
)

var (
	ErrTankErrorGroupNotFound = errors.New("didn't found tank group")
	ErrTankErrorServerError   = func(errorMsg string) error {
		return fmt.Errorf("server error: %s", errorMsg)
	}
)
