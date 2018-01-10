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

func ManagerRouterScan(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetScan(w, r)
	case "POST":
		RequestPostScan(w, r)
	case "UPDATE":
		RequestUpdateScan(w, r)
	case "DELETE":
		RequestDeleteScan(w, r)
	default:
	}
}

func RequestGetScan(w http.ResponseWriter, r *http.Request) {
	var data []model.ScanModel

	db, err := database.NewMySQLDBScan()
	scan, err := db.ListScan()

	if err != nil {
		fmt.Println("error")
	}

	for _, value := range scan {
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

func RequestPostScan(w http.ResponseWriter, r *http.Request) {

	var data model.ScanModel
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

	db, err := database.NewMySQLDBScan()
	if err != nil {
		fmt.Println(err)
	}
	idscan, err := db.AddScan(&data)

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
		fmt.Println("New ID_Scan: ", idscan)
		response = model.ResponseRequestIdModel{
			Code: 200,
			Msg:  "Datos guardados",
			ID:   idscan,
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

func RequestUpdateScan(w http.ResponseWriter, r *http.Request) {

	var data model.ScanModel
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

	db, err := database.NewMySQLDBScan()
	if err != nil {
		fmt.Println(err)
	}
	err = db.UpdateScan(&data)

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

func RequestDeleteScan(w http.ResponseWriter, r *http.Request) {
	var data model.ScanModel
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

	db, err := database.NewMySQLDBScan()
	if err != nil {
		fmt.Println(err)
	}

	err = db.DeleteScan(data.ID)

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
