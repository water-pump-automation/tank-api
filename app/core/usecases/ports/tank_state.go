package ports

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"water-tank-api/app/core/entity/water_tank"
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

func MapTankStateEnum(tankState water_tank.State) string {
	switch tankState {
	case water_tank.Empty:
		return "EMPTY"
	case water_tank.Filling:
		return "FILLING"
	case water_tank.Full:
		return "FULL"
	default:
		return "UNKOWN"
	}
}

func ConvertCapacityToLiters(level water_tank.Capacity) string {
	return fmt.Sprintf("%1.2fL", level)
}

func ConverLitersToCapacity(liters string) water_tank.Capacity {
	substrings := strings.Split(liters, "L")
	capacity, _ := strconv.ParseFloat(substrings[0], 32)

	return water_tank.Capacity(capacity)
}
