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

func ManagerRouterUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetUser(w, r)
	case "POST":
		RequestPostUser(w, r)
	case "PUT":
		RequestUpdateUser(w, r)
	case "DELETE":
		RequestDeleteUser(w, r)
	default:
	}
}

func RequestGetUser(w http.ResponseWriter, r *http.Request) {
	var data []model.UserModel

	db, err := database.NewMySQLDBUser()
	user, err := db.ListUser()

	if err != nil {
		fmt.Println("error")
	}

	for _, value := range user {
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

func RequestPostUser(w http.ResponseWriter, r *http.Request) {

	var data model.UserModel
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
		fmt.Printf("error json:")
		fmt.Println(err)
	}

	db, err := database.NewMySQLDBUser()
	if err != nil {
		fmt.Println(err)
	}

	if data.Name != "" {
		//registro de usuario
		iduser, err := db.AddUser(&data)
		fmt.Printf("Registro completo, id: %d", iduser)

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
			response = model.ResponseRequestIdModel{
				Code: 200,
				Msg:  "Registro completo",
				ID:   0,
			}
		}

	} else {
		//login de usuario
		user, err := db.GetUser(data.Email)
		fmt.Printf("Login completo, name: %s", user.Name)

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
			response = model.ResponseRequestIdModel{
				Code: 200,
				Msg:  "Login completo",
				ID:   0,
			}
		}

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

func RequestUpdateUser(w http.ResponseWriter, r *http.Request) {

	var data model.UserModel
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

	db, err := database.NewMySQLDBUser()
	if err != nil {
		fmt.Println(err)
	}
	err = db.UpdateUser(&data)

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

func RequestDeleteUser(w http.ResponseWriter, r *http.Request) {
	var data model.UserModel
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

	db, err := database.NewMySQLDBUser()
	if err != nil {
		fmt.Println(err)
	}

	err = db.DeleteUser(data.ID)

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
