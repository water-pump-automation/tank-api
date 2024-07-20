package create_tank

import (
	"errors"
	"fmt"
)

var (
	ErrWaterTankAlreadyExists       = errors.New("water tank already exists")
	ErrWaterTankInvalidName         = errors.New("invalid tank name")
	ErrWaterTankInvalidGroup        = errors.New("invalid tank group")
	ErrWaterTankMaximumCapacityZero = errors.New("invalid capacity. Must be bigger than 0")
	ErrWaterTankErrorServerError    = func(errorMsg string) error {
		return fmt.Errorf("server error: %s", errorMsg)
	}
)
