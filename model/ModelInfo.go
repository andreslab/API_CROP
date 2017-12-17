package model

type WeatherModel struct{
	state string `json:"state"` //estado: nublado, soleado, etc
	temperature string `json:"temp"`
	wet string `json:"wet"` //humedad
	windSpeed string `json:"wind"`
	precipitationProbability string `json:"pp"`
}

type ZoneModel struct{
	city string `json:"city"`
	latitude string `json:"lat"`
	longitude string `json:"longitude"`
	address string `json:"address"`
}

type RecorModel struct{
	dateCreated string `json:"created"`
	dateEdit string `json:"edit"`
	note string `json:"note"`
	title string `json:"title"`
	tag string `json:"tag"`
	alarm string `json:"alarm"`
}

type IrrigationModel struct{
	alarm string `json:"alarm"`
	quantity string `json:"quantity"`
	phLevel string `json:"ph"`
}

type NutrientModel struct{
	component string `json:"component"`
	advantage string `json:"advantage"`
	atomicNumber string `json:"atom"`
}

type PestsModel struct{
	category string `json:"type"`
	causes string `json:"causes"`
	consequences string `json:"consequences"`
	solution string `json:"solution"`
}