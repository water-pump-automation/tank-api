package get_tank

import (
	"time"
	"water-tank-api/core/entity/data"
)

type WaterTankState struct {
	Tank     *data.WaterTankState `json:"tank"`
	Datetime time.Time            `json:"datetime"`
}
