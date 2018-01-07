package model

type scanModel struct {
	ID         int64
	Zone       string `json:"zone"`
	Distance   string `json:"distance"`
	ZoneWidth  string `json:"zone_width"`
	ZoneHeight string `json:"zone_height"`
	IDUser     int64  `json:"id_user"`
}
