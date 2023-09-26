package update_tank_state

import (
	"errors"
	"fmt"
)

var (
	WaterTankCurrentWaterLevelSmallerThanZero = errors.New("Invalid water level. Smaller than 0")
	WaterTankCurrentWaterLevelBiggerThanMax   = errors.New("Invalid water level. Bigger than maximum capacity")
	WaterTankInvalidToken                     = errors.New("Invalid access token!")
	WaterTankErrorNotFound                    = func(tank string) error {
		return fmt.Errorf("Didn't found %s tank", tank)
	}
	WaterTankErrorServerError = func(errorMsg string) error {
		return fmt.Errorf("Server error: %s", errorMsg)
	}
)
