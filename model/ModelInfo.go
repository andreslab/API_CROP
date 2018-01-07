package model

//información que el usuario provee a la aplciación

type infoModel struct {
	Crop       string        `json:"crop"`
	SowingType string        `json:"sowing"` //sowing: siembra
	Zona       infoZoneModel `json:"zone"`
}

type infoZoneModel struct {
	City      string `json:"city"`
	Latitude  string `json:"lat"`
	Longitude string `json:"longitude"`
	Address   string `json:"address"`
}
