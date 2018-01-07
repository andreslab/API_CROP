package model

type actionCreateNoteModel struct {
	DateCreated string `json:"created"`
	DateEdit    string `json:"edit"`
	Body        string `json:"body"`
	Title       string `json:"title"`
	Tag         string `json:"tag"`
	Alarm       string `json:"alarm"`
}

type actionCreateIrrigationAlarmModel struct {
	DateCreated string `json:"alarm"`
	DateAlarm   string `json:"date_alarm"`
	Quantity    string `json:"quantity"`
	PhLevel     string `json:"ph"`
}
