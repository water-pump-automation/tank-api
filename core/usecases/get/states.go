package get

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	data "water-tank-api/core/entity/water_tank"
)

type WaterTankState struct {
	Name              string     `json:"name"`
	Group             string     `json:"group"`
	MaximumCapacity   string     `json:"maximum_capacity"`
	TankState         string     `json:"tank_state"`
	CurrentWaterLevel string     `json:"current_water_level"`
	Datetime          *time.Time `json:"datetime,omitempty"`
}

type WaterTankGroupState struct {
	Tanks    []*WaterTankState `json:"tanks"`
	Datetime time.Time         `json:"datetime"`
}

func MapTankStateEnum(tankState data.State) string {
	switch tankState {
	case data.Empty:
		return "EMPTY"
	case data.Filling:
		return "FILLING"
	case data.Full:
		return "FULL"
	default:
		return "UNKOWN"
	}
}

func ConvertCapacityToLiters(level data.Capacity) string {
	return fmt.Sprintf("%1.2fL", level)
}

func ConverLitersToCapacity(liters string) data.Capacity {
	substrings := strings.Split(liters, "L")
	capacity, _ := strconv.ParseFloat(substrings[0], 32)

	return data.Capacity(capacity)
}
