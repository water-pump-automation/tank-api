package get_tank

import (
	"errors"
	"fmt"
)

var (
	ErrWaterTankErrorNotFound    = errors.New("didn't found tank")
	ErrWaterTankErrorServerError = func(errorMsg string) error {
		return fmt.Errorf("server error: %s", errorMsg)
	}
)
