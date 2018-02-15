package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreslab/prj_api_crop/model"
)

func ManagerRouterStorage(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetStorage(w, r)
	case "POST":
		fmt.Println("Resquest POST not wonrking")
	case "UPDATE":
		fmt.Println("Resquest UPDATE not wonrking")
	case "DELETE":
		fmt.Println("Resquest DELETE not wonrking")
		//RequestDeleteStorage(w, r)
	default:
	}
}

func RequestGetStorage(w http.ResponseWriter, r *http.Request) {
	var data model.StorageModel

	/*db, err := database.NewMySQLDBStorage()
	Storage, err := db.ListStorage()

	if err != nil {
		fmt.Println("error")
	}

	for _, value := range Storage {
		data = append(data, *value)
	}*/

	data = model.StorageModel{
		Data: []model.ContentAbstractAction{
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Fertilizantes para cuidado",
				Description: "",
				Type:        "note",
				DateCreated: "12/01/18",
				DateEdit:    "",
				IDUser:      14,
			},
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Siembra",
				Description: "",
				Type:        "alarm",
				DateCreated: "02/02/18",
				DateEdit:    "",
				IDUser:      14,
			},
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Revisión del cultivo",
				Description: "",
				Type:        "note",
				DateCreated: "20/01/18",
				DateEdit:    "",
				IDUser:      14,
			},
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Riego de amanecer",
				Description: "",
				Type:        "photo",
				DateCreated: "06/07/17",
				DateEdit:    "",
				IDUser:      14,
			},
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Riego de amanecer",
				Description: "",
				Type:        "alarm",
				DateCreated: "06/07/17",
				DateEdit:    "",
				IDUser:      14,
			},
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Riego de amanecer",
				Description: "",
				Type:        "note",
				DateCreated: "06/07/17",
				DateEdit:    "",
				IDUser:      14,
			},
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Riego de amanecer",
				Description: "",
				Type:        "alarm",
				DateCreated: "06/07/17",
				DateEdit:    "",
				IDUser:      14,
			},
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Riego de amanecer",
				Description: "",
				Type:        "photo",
				DateCreated: "06/07/17",
				DateEdit:    "",
				IDUser:      14,
			},
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Riego de amanecer",
				Description: "",
				Type:        "note",
				DateCreated: "06/07/17",
				DateEdit:    "",
				IDUser:      14,
			},
			model.ContentAbstractAction{
				ID:          0,
				Title:       "Riego de amanecer",
				Description: "",
				Type:        "alarm",
				DateCreated: "06/07/17",
				DateEdit:    "",
				IDUser:      14,
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
