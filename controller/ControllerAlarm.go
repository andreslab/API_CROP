package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andreslab/prj_api_crop/database"
	"github.com/andreslab/prj_api_crop/model"
)

func ManagerRouterAlarm(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetAlarm(w, r)
	case "POST":
		RequestPostAlarm(w, r)
	case "UPDATE":
		RequestUpdateAlarm(w, r)
	case "DELETE":
		RequestDeleteAlarm(w, r)
	default:
	}
}

func RequestGetAlarm(w http.ResponseWriter, r *http.Request) {
	var data []model.ActionCreateAlarmModel

	db, err := database.NewMySQLDBAlarm()
	alarm, err := db.ListAlarm()

	if err != nil {
		fmt.Println("error")
	}

	for _, value := range alarm {
		data = append(data, *value)
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

func RequestPostAlarm(w http.ResponseWriter, r *http.Request) {

	var data model.ActionCreateAlarmModel
	var response model.ResponseRequestIdModel

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	dataJSON := buf.String()
	fmt.Println("Buffer", dataJSON)
	/*byteArray := []byte(dataJSON)
	fmt.Println("StringTobytes: ", byteArray)
	dataJSON := string(byteArray)
	fmt.Println("dataJson:", dataJSON)*/

	err := json.Unmarshal([]byte(dataJSON), &data)
	if err != nil {
		fmt.Println(err)
	}

	db, err := database.NewMySQLDBAlarm()
	if err != nil {
		fmt.Println(err)
	}
	idalarm, err := db.AddAlarm(&data)

	if err != nil {
		log.Fatal(err)
		log.Println("json.Compact:", err)
		if serr, ok := err.(*json.SyntaxError); ok {
			log.Println("Occurred at offset:", serr.Offset)
		}
		response = model.ResponseRequestIdModel{
			Code: 500,
			Msg:  "Error al guardar los datos",
			ID:   0,
		}
		//header
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
	} else {
		fmt.Println("New ID_Alarm: ", idalarm)
		response = model.ResponseRequestIdModel{
			Code: 200,
			Msg:  "Datos guardados",
			ID:   idalarm,
		}
		//header
		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

}

func RequestUpdateAlarm(w http.ResponseWriter, r *http.Request) {

	var data model.ActionCreateAlarmModel
	var response model.ResponseRequestModel

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	dataJSON := buf.String()
	fmt.Println("Buffer", dataJSON)
	/*byteArray := []byte(dataJSON)
	fmt.Println("StringTobytes: ", byteArray)
	dataJSON := string(byteArray)
	fmt.Println("dataJson:", dataJSON)*/

	err := json.Unmarshal([]byte(dataJSON), &data)
	if err != nil {
		fmt.Println(err)
	}

	db, err := database.NewMySQLDBAlarm()
	if err != nil {
		fmt.Println(err)
	}
	err = db.UpdateAlarm(&data)

	if err != nil {
		log.Fatal(err)
		log.Println("json.Compact:", err)
		if serr, ok := err.(*json.SyntaxError); ok {
			log.Println("Occurred at offset:", serr.Offset)
		}
		response = model.ResponseRequestModel{
			Code: 500,
			Msg:  "Error al actualizar los datos",
		}
		//header
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
	} else {
		response = model.ResponseRequestModel{
			Code: 200,
			Msg:  "Datos actualizados correctamente",
		}
		//header
		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}

func RequestDeleteAlarm(w http.ResponseWriter, r *http.Request) {
	var data model.ActionCreateAlarmModel
	var response model.ResponseRequestModel

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	dataJSON := buf.String()
	fmt.Println("Buffer", dataJSON)
	/*byteArray := []byte(dataJSON)
	fmt.Println("StringTobytes: ", byteArray)
	dataJSON := string(byteArray)
	fmt.Println("dataJson:", dataJSON)*/

	err := json.Unmarshal([]byte(dataJSON), &data)
	if err != nil {
		fmt.Println(err)
	}

	db, err := database.NewMySQLDBAlarm()
	if err != nil {
		fmt.Println(err)
	}

	err = db.DeleteAlarm(data.ID)

	if err != nil {
		log.Fatal(err)
		log.Println("json.Compact:", err)
		if serr, ok := err.(*json.SyntaxError); ok {
			log.Println("Occurred at offset:", serr.Offset)
		}
		response = model.ResponseRequestModel{
			Code: 500,
			Msg:  "Error al actualizar los datos",
		}
		//header
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
	} else {
		response = model.ResponseRequestModel{
			Code: 200,
			Msg:  "Datos actualizados correctamente",
		}
		//header
		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}
