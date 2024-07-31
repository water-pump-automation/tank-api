package create_tank

import (
	"errors"
	"fmt"
)

var (
	ErrTankAlreadyExists       = errors.New(" tank already exists")
	ErrTankInvalidName         = errors.New("invalid tank name")
	ErrTankInvalidGroup        = errors.New("invalid tank group")
	ErrTankMaximumCapacityZero = errors.New("invalid capacity. Must be bigger than 0")
	ErrTankErrorServerError    = func(errorMsg string) error {
		return fmt.Errorf("server error: %s", errorMsg)
	}
)
