package model

type ScanModel struct {
	ID         int64
	Created    string `json:"created"`
	Distance   string `json:"distance"`
	ZoneWidth  string `json:"zone_width"`
	ZoneHeight string `json:"zone_height"`
	Latitude   string `json:"lat"`
	Longitude  string `json:"lomg"`
	City       string `json:"city"`
	IDUser     int64  `json:"id_user"`
}
