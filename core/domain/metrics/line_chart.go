package metrics

import "time"

// Point represents a single data point with X and Y coordinates
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// TimePoint represents a data point with a timestamp
type TimePoint struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

// Line represents a single line in the chart
type Line struct {
	Name       string       `json:"name"`
	Points     *[]Point     `json:"points,omitempty"`
	TimePoints *[]TimePoint `json:"timePoints,omitempty"`
}

// Axis represents chart axis configuration
type Axis struct {
	Label string  `json:"label"`
	Min   float64 `json:"min,omitempty"`
	Max   float64 `json:"max,omitempty"`
}

type TemporalAxis struct {
	Label string    `json:"label"`
	Min   time.Time `json:"min,omitempty"`
	Max   time.Time `json:"max,omitempty"`
}

// LineChart represents the complete multi-line chart
type LineChart struct {
	Lines        []Line        `json:"lines"`
	XAxis        *Axis         `json:"xAxis"`
	YAxis        Axis          `json:"yAxis"`
	TemporalAxis *TemporalAxis `json:"temporalAxis,omitempty"`
}
