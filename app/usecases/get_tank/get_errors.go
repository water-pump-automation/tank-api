package get_tank

import (
	"errors"
	"fmt"
)

var (
	ErrTankErrorNotFound    = errors.New("didn't found tank")
	ErrTankErrorServerError = func(errorMsg string) error {
		return fmt.Errorf("server error: %s", errorMsg)
	}
)
