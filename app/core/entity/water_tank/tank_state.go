package water_tank

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type WaterTankState struct {
	Name              string     `json:"name"`
	Group             string     `json:"group"`
	MaximumCapacity   string     `json:"maximum_capacity"`
	TankState         string     `json:"tank_state"`
	CurrentWaterLevel string     `json:"current_water_level"`
	LastFullTime      time.Time  `json:"last_full_time"`
	Datetime          *time.Time `json:"datetime,omitempty"`
}

type WaterTankGroupState struct {
	Tanks    []*WaterTankState `json:"tanks"`
	Datetime time.Time         `json:"datetime"`
}

func MapTankStateEnum(tankState State) string {
	switch tankState {
	case Empty:
		return "EMPTY"
	case Filling:
		return "FILLING"
	case Full:
		return "FULL"
	default:
		return "UNKOWN"
	}
}

func ConvertCapacityToLiters(level Capacity) string {
	return fmt.Sprintf("%1.2fL", level)
}

func ConverLitersToCapacity(liters string) Capacity {
	substrings := strings.Split(liters, "L")
	capacity, _ := strconv.ParseFloat(substrings[0], 32)

	return Capacity(capacity)
}
