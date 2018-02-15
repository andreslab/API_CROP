package model

type StorageModel struct {
	/*DataNote  []ActionCreateNoteModel  `json:"data_note"`
	DataAlarm []ActionCreateAlarmModel `json:"data_alarm"`
	DataPhoto []ActionCreatePhotoModel `json:"data_photo"`*/
	Data []ContentAbstractAction `json:"data"`
}
