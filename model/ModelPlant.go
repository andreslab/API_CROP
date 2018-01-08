package model

type PlantModel struct {
	ID             int64
	Name           string                    `json:"name"`
	Kingdom        string                    `json:"kingdom"`
	SeasonEnv      string                    `json:"season"`
	TemperatureEnv string                    `json:"temp"`
	Nutrient       PlantNutrientFormulaModel `json:"nutrient"`
	Pests          PlantPestsModel           `json:"pests"`
}

type PlantNutrientFormulaModel struct {
	ID        int64
	Formula   string `json:"formula"`
	Advantage string `json:"advantage"`
}

type PlantPestsModel struct {
	ID           int64
	Category     string `json:"type"`
	Causes       string `json:"causes"`
	Consequences string `json:"consequences"`
	Solution     string `json:"solution"`
}
