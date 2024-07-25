package get_group

import (
	"errors"
	"fmt"
)

var (
	ErrWaterTankErrorGroupNotFound = errors.New("didn't found tank group")
	ErrWaterTankErrorServerError   = func(errorMsg string) error {
		return fmt.Errorf("server error: %s", errorMsg)
	}
)
