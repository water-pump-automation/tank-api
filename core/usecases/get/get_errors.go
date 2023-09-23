package get_tank

import (
	"fmt"
)

var (
	WaterTankErrorNotFound = func(tank string) error {
		return fmt.Errorf("Didn't found %s tank", tank)
	}
	WaterTankErrorServerError = func(errorMsg string) error {
		return fmt.Errorf("Server error: %s", errorMsg)
	}
)
