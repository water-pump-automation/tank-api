package update_tank_state

import (
	"errors"
	"fmt"
)

var (
	ErrTankCurrentLevelSmallerThanZero = errors.New("invalid  level. Smaller than 0")
	ErrTankCurrentLevelBiggerThanMax   = errors.New("invalid  level. Bigger than maximum capacity")
	ErrTankErrorNotFound               = errors.New("didn't found tank")
	ErrTankErrorServerError            = func(errorMsg string) error {
		return fmt.Errorf("server error: %s", errorMsg)
	}
)
