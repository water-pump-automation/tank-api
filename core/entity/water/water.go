package water

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
	RegisterWaterStats(waterTankName string, stats *Water) (err error)
	UpdateWaterStats(waterTankName string, stats *Water) (err error)
	GetWaterStats(waterTankName string) (state *Water, err error)
}
