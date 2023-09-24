package get_tank

import (
	"time"
	data "water-tank-api/core/entity/water_tank"
)

type WaterTankState struct {
	Tank     *data.WaterTankState `json:"tank"`
	Datetime time.Time            `json:"datetime"`
}
