package register_tank

import (
	"errors"
	"fmt"
)

var (
	WaterTankAlreadyExists       = errors.New("Water tank already exists")
	WaterTankInvalidName         = errors.New("Invalid tank name")
	WaterTankInvalidGroup        = errors.New("Invalid tank group")
	WaterTankMaximumCapacityZero = errors.New("Invalid capacity. Must be bigger than 0")
	WaterTankErrorServerError    = func(errorMsg string) error {
		return fmt.Errorf("Server error: %s", errorMsg)
	}
)
