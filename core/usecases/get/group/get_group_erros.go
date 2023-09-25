package group

import (
	"fmt"
)

var (
	WaterTankErrorGroupNotFound = func(group string) error {
		return fmt.Errorf("Didn't found %s tank group", group)
	}
	WaterTankErrorServerError = func(errorMsg string) error {
		return fmt.Errorf("Server error: %s", errorMsg)
	}
)
