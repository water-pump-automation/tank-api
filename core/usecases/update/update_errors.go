package update_tank

import (
	"errors"
	"fmt"
)

var (
	WaterTankCurrentWaterLevelSmallerThanZero = errors.New("Invalid water level. Smaller than 0")
	WaterTankCurrentWaterLevelBiggerThanMax   = errors.New("Invalid water level. Bigger than maximum capacity")
	WaterTankErrorServerError                 = func(errorMsg string) error {
		return fmt.Errorf("Server error: %s", errorMsg)
	}
)
