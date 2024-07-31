package ports

import (
	"fmt"
	"strconv"
	"strings"
	"tank-api/app/entity/tank"
	"time"
)

const (
	EMPTY   = "EMPTY"
	FILLING = "FILLING"
	FULL    = "FULL"
	UNKOWN  = "UNKOWN"
)

type TankState struct {
	Name            string     `json:"name"`
	Group           string     `json:"group"`
	MaximumCapacity string     `json:"maximum_capacity"`
	TankState       string     `json:"tank_state"`
	CurrentLevel    string     `json:"current_level"`
	LastFullTime    *time.Time `json:"last_full_time,omitempty"`
	Datetime        time.Time  `json:"datetime,omitempty"`
}

type TankGroupState struct {
	Tanks    []*TankState `json:"tanks"`
	Datetime time.Time    `json:"datetime"`
}

func MapState(level tank.Capacity, maximumCapacity tank.Capacity) tank.State {
	if level == maximumCapacity {
		return tank.Full
	} else if level == 0 {
		return tank.Empty
	} else if level < maximumCapacity {
		return tank.Filling
	}
	return tank.Invalid
}

func ConvertState(tankState tank.State) string {
	switch tankState {
	case tank.Empty:
		return EMPTY
	case tank.Filling:
		return FILLING
	case tank.Full:
		return FULL
	default:
		return UNKOWN
	}
}

func ConvertCapacityToLiters(level tank.Capacity) string {
	return fmt.Sprintf("%1.2fL", level)
}

func ConverLitersToCapacity(liters string) tank.Capacity {
	substrings := strings.Split(liters, "L")
	capacity, _ := strconv.ParseFloat(substrings[0], 32)

	return tank.Capacity(capacity)
}
