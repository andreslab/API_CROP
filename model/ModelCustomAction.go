package model

type ContentAbstractAction struct {
	ID          int64
	Title       string `json:"title"`
	Description string `json:"body"`
	Type        string `json:"type"`
	DateCreated string `json:"created"`
	DateEdit    string `json:"edited"`
	IDUser      int64  `json:"iduser"`
}

type ActionCreateNoteModel struct {
	ID          int64
	Title       string `json:"title"`
	Body        string `json:"body"`
	Tag         string `json:"tag"`
	Alarm       string `json:"alarm"`
	DateCreated string `json:"created"`
	DateEdit    string `json:"edit"`
	IDUser      int64  `json:"iduser"`
}

type ActionCreateAlarmModel struct {
	ID          int64
	Title       string `json:"title"`
	DateCreated string `json:"created"`
	DateAlarm   string `json:"alarm"`
	IDUser      int64  `json:"iduser"`
}

type ActionCreatePhotoModel struct {
	ID          int64
	Title       string `json:"title"`
	DateCreated string `json:"created"`
	IDPhoto     string `json:"id_photo"`
	IDUser      int64  `json:"iduser"`
}

/*
ittigación
Quantity    string `json:"quantity"`
PhLevel     string `json:"ph"`
*/
