package model

type plantModel struct {
	ID             int64
	Name           string                    `json:"name"`
	Kingdom        string                    `json:"kingdom"`
	SeasonEnv      string                    `json:"season"`
	TemperatureEnv string                    `json:"temp"`
	Nutrient       plantNutrientFormulaModel `json:"nutrient"`
	Pests          plantPestsModel           `json:"pests"`
}

type plantNutrientFormulaModel struct {
	ID        int64
	Formula   string `json:"formula"`
	Advantage string `json:"advantage"`
}

type plantPestsModel struct {
	ID           int64
	Category     string `json:"type"`
	Causes       string `json:"causes"`
	Consequences string `json:"consequences"`
	Solution     string `json:"solution"`
}
