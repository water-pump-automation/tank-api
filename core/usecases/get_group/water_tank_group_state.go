package get_group

import (
	"time"
	"water-tank-api/core/entity/data"
)

type WaterTankGroupState struct {
	Group    string                 `json:"group"`
	Tank     []*data.WaterTankState `json:"tanks"`
	Datetime time.Time              `json:"datetime"`
}
