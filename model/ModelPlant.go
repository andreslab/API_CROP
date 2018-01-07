package model

type plantModel struct {
	ID       int64
	Name     string                      `json:"name"`
	Kingdom  string                      `json:"kingdom"`
	Nutrient plantNutrientModel          `json:"nutrient"`
	Weather  plantWeatherEnviromentModel `josn:"weather"`
	Pests    plantPestsModel             `json:"pests"`
}

type plantNutrientModel struct {
	Component    string `json:"component"`
	Advantage    string `json:"advantage"`
	AtomicNumber string `json:"atom"`
}

type plantWeatherEnviromentModel struct {
	State                    string `json:"state"` //estado: nublado, soleado, etc
	Temperature              string `json:"temp"`
	Wet                      string `json:"wet"` //humedad
	WindSpeed                string `json:"wind"`
	PrecipitationProbability string `json:"pp"`
}

type plantPestsModel struct {
	Category     string `json:"type"`
	Causes       string `json:"causes"`
	Consequences string `json:"consequences"`
	Solution     string `json:"solution"`
}
