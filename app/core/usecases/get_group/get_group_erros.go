package get_group

import (
	"errors"
	"fmt"
)

var (
	ErrWaterTankErrorGroupNotFound = func(group string) error {
		return fmt.Errorf("didn't found %s tank group", group)
	}
	ErrWaterTankMissingGroup     = errors.New("missing tank group")
	ErrWaterTankErrorServerError = func(errorMsg string) error {
		return fmt.Errorf("server error: %s", errorMsg)
	}
)
