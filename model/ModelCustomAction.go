package model

type actionCreateNoteModel struct {
	ID          int64
	Title       string `json:"title"`
	Body        string `json:"body"`
	Tag         string `json:"tag"`
	Alarm       string `json:"alarm"`
	DateCreated string `json:"created"`
	DateEdit    string `json:"edit"`
	IDUser      int64  `json:"iduser"`
}

type actionCreateAlarmModel struct {
	Title       string `json:"title"`
	DateCreated string `json:"created"`
	DateAlarm   string `json:"alarm"`
	IDUser      int64  `json:"iduser"`
}

/*
ittigación
Quantity    string `json:"quantity"`
PhLevel     string `json:"ph"`
*/
