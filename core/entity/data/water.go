package data

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
