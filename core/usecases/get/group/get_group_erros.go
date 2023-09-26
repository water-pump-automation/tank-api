package group

import (
	"errors"
	"fmt"
)

var (
	WaterTankErrorGroupNotFound = func(group string) error {
		return fmt.Errorf("Didn't found %s tank group", group)
	}
	WaterTankMissingGroup     = errors.New("Missing tank group!")
	WaterTankErrorServerError = func(errorMsg string) error {
		return fmt.Errorf("Server error: %s", errorMsg)
	}
)
