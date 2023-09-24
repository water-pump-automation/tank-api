package get_group

import (
	"time"
	data "water-tank-api/core/entity/water_tank"
)

type WaterTankGroupState struct {
	Group    string                 `json:"group"`
	Tank     []*data.WaterTankState `json:"tanks"`
	Datetime time.Time              `json:"datetime"`
}
