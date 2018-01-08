package model

//información que el usuario provee a la aplciación

type infoModel struct {
	ID         int64
	Created    string `json:"created"`
	Crop       string `json:"crop"`
	SowingType string `json:"sowing"` //sowing: siembra
	City       string `json:"city"`
	Address    string `json:"address"`
	IDUser     int64  `json:"iduser"`
}
