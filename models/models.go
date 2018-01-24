package models

// Region of the spot
type Region struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name:omitempty"`
}

// Spot (beach)
type Spot struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name:omitempty"`
	RegionID string `json:"region:omitempty"`
	// Region *Region `json:"region:omitempty"`
	// Winds    []int  `json:"winds:omitempty"`
}
