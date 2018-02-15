package model

type ProfileModel struct {
	ID      int64
	Title   string `json:"title"`
	Content string `json:"content"`
}
