package update_tank_state

import (
	"errors"
	"fmt"
)

var (
	ErrWaterTankCurrentWaterLevelSmallerThanZero = errors.New("invalid water level. Smaller than 0")
	ErrWaterTankCurrentWaterLevelBiggerThanMax   = errors.New("invalid water level. Bigger than maximum capacity")
	ErrWaterTankErrorNotFound                    = errors.New("didn't found tank")
	ErrWaterTankErrorServerError                 = func(errorMsg string) error {
		return fmt.Errorf("server error: %s", errorMsg)
	}
)
