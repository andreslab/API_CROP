package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreslab/prj_api_crop/model"
)

func ManagerRouterProfile(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetProfile(w, r)
	case "POST":
		fmt.Println("Resquest POST not wonrking")
	case "UPDATE":
		fmt.Println("Resquest UPDATE not wonrking")
	case "DELETE":
		fmt.Println("Resquest DELETE not wonrking")
		//RequestDeleteProfile(w, r)
	default:
	}
}

func RequestGetProfile(w http.ResponseWriter, r *http.Request) {
	var data []model.ProfileModel

	/*db, err := database.NewMySQLDBProfile()
	Profile, err := db.ListProfile()
	if err != nil {
		fmt.Println("error")
	}
	for _, value := range Profile {
		data = append(data, *value)
	}*/

	data = []model.ProfileModel{
		model.ProfileModel{
			Title:   "Estado del cultivo",
			Content: "En riesgo",
		},
		model.ProfileModel{
			Title:   "Num. escaner",
			Content: "10",
		},
		model.ProfileModel{
			Title:   "Ultimo cultivo escaneado",
			Content: "Tomate",
		},
		model.ProfileModel{
			Title:   "Total de zona",
			Content: "100 ha",
		},
		model.ProfileModel{
			Title:   "Plaga frecuente",
			Content: "viruela",
		},
		model.ProfileModel{
			Title:   "Número de plagas registradas",
			Content: "5",
		},
		model.ProfileModel{
			Title:   "Número de cultvos",
			Content: "2",
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
