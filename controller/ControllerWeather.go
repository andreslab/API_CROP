package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreslab/prj_api_crop/model"
)

func ManagerRouterWeather(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetWeather(w, r)
	case "POST":
		fmt.Println("Resquest POST not wonrking")
	case "UPDATE":
		fmt.Println("Resquest UPDATE not wonrking")
	case "DELETE":
		fmt.Println("Resquest DELETE not wonrking")
		//RequestDeleteWeather(w, r)
	default:
	}
}

func RequestGetWeather(w http.ResponseWriter, r *http.Request) {
	var data model.ContentWeatherModel

	/*db, err := database.NewMySQLDBWeather()
	Weather, err := db.ListWeather()

	if err != nil {
		fmt.Println("error")
	}

	for _, value := range Weather {
		data = append(data, *value)
	}*/

	data = model.ContentWeatherModel{
		Data: []model.WeatherModel{
			model.WeatherModel{
				Day:                      "14/01/18",
				State:                    "nublado",
				Temperature:              "25",
				Wet:                      "10",
				WindSpeed:                "5",
				PrecipitationProbability: "5",
			},
			model.WeatherModel{
				Day:                      "14/01/18",
				State:                    "nublado",
				Temperature:              "25",
				Wet:                      "10",
				WindSpeed:                "5",
				PrecipitationProbability: "5",
			},
			model.WeatherModel{
				Day:                      "14/01/18",
				State:                    "nublado",
				Temperature:              "25",
				Wet:                      "10",
				WindSpeed:                "5",
				PrecipitationProbability: "5",
			},
			model.WeatherModel{
				Day:                      "14/01/18",
				State:                    "nublado",
				Temperature:              "25",
				Wet:                      "10",
				WindSpeed:                "5",
				PrecipitationProbability: "5",
			},
			model.WeatherModel{
				Day:                      "14/01/18",
				State:                    "nublado",
				Temperature:              "25",
				Wet:                      "10",
				WindSpeed:                "5",
				PrecipitationProbability: "5",
			},
		},
	}

	//header
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
