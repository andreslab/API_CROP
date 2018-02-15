package model

type ContentWeatherModel struct {
	Data []WeatherModel `json:"data"`
}

type WeatherModel struct {
	Day                      string `json:"day"`
	State                    string `json:"state"` //estado: nublado, soleado, etc
	Temperature              string `json:"temp"`
	Wet                      string `json:"wet"` //humedad
	WindSpeed                string `json:"wind"`
	PrecipitationProbability string `json:"pp"`
}
