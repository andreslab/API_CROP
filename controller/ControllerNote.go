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

func ManagerRouterNote(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetNote(w, r)
	case "POST":
		RequestPostNote(w, r)
	case "UPDATE":
		RequestUpdateNote(w, r)
	case "DELETE":
		RequestDeleteNote(w, r)
	default:
	}
}

func RequestGetNote(w http.ResponseWriter, r *http.Request) {
	var data []model.ActionCreateNoteModel

	db, err := database.NewMySQLDBNote()
	note, err := db.ListNote()

	if err != nil {
		fmt.Println("error")
	}

	for _, value := range note {
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

func RequestPostNote(w http.ResponseWriter, r *http.Request) {
	var data model.ActionCreateNoteModel
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

	db, err := database.NewMySQLDBNote()
	if err != nil {
		fmt.Println(err)
	}
	idnote, err := db.AddNote(&data)

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
		fmt.Println("New ID_Note: ", idnote)
		response = model.ResponseRequestIdModel{
			Code: 200,
			Msg:  "Datos guardados",
			ID:   idnote,
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

func RequestUpdateNote(w http.ResponseWriter, r *http.Request) {

	var data model.ActionCreateNoteModel
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

	db, err := database.NewMySQLDBNote()
	if err != nil {
		fmt.Println(err)
	}
	err = db.UpdateNote(&data)

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

func RequestDeleteNote(w http.ResponseWriter, r *http.Request) {
	var data model.ActionCreateNoteModel
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

	db, err := database.NewMySQLDBNote()
	if err != nil {
		fmt.Println(err)
	}

	err = db.DeleteNote(data.ID)

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
