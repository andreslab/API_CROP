package model

type ScanModel struct {
	zone        string  `json:"zone"`
	distance    string  `json:"distance"`
	zone_width  float32 `json:"zone_width"`
	zone_height float32 `json:"zone_height"`
}
