package model

type PestsModel struct {
	ID          int64
	Name        string `json:"name"`
	Description string `json:"description"`
	Treatment   string `json:"treatment"`
	Percent     string `json:"percent"`
}
