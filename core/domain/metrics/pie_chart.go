package metrics

// PieSlice represents a single slice of the pie chart
type PieSlice struct {
	Label      string  `json:"label"`
	Value      float64 `json:"value"`
	Percentage float64 `json:"percentage"`
}

// PieChart represents the complete pie chart data
type PieChart struct {
	Slices []PieSlice `json:"slices"`
	Total  float64    `json:"total"`
}
