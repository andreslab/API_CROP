package model

type weatherModel struct {
	State                    string `json:"state"` //estado: nublado, soleado, etc
	Temperature              string `json:"temp"`
	Wet                      string `json:"wet"` //humedad
	WindSpeed                string `json:"wind"`
	PrecipitationProbability string `json:"pp"`
}
