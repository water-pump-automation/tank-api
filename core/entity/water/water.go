package water

import stack "water-tank-api/core/entity/error_stack"

type WaterQuality int

type Temperature string

const (
	Ok WaterQuality = iota
	Warning
	Inappropriate
)

type Water struct {
	Quality     WaterQuality
	Reason      string
	Temperature Temperature
}

type WaterData interface {
	RegisterWaterStats(waterTankName string, stats *Water) (err stack.ErrorStack)
	UpdateWaterStats(waterTankName string, stats *Water) (err stack.ErrorStack)
	GetWaterStats(waterTankName string) (state *Water, err stack.ErrorStack)
}
