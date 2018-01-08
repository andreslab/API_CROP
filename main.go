package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/andreslab/prj_api_crop/database"
)

func main() {

	/*err := database.CreateTableUser()
	if err != nil {
		fmt.Println(err)
	}*/
	/*conf := &database.MySQLConfig{
		Username:   utils.User,
		Password:   utils.Pass,
		Host:       utils.Host,
		Port:       utils.Port,
		UnixSocket: "linux",
	}*/

	db, err := database.NewMySQLDB()
	user, err := db.ListUser()

	fmt.Println("users")
	fmt.Println(user[0].Name)

	if err != nil {
		fmt.Println("error")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", CheckState)
	mux.HandleFunc("/crop/alarm")
	mux.HandleFunc("/crop/note", CheckState)
	mux.HandleFunc("/crop/result", CheckState)
	mux.HandleFunc("/crop/scan", CheckState)
	mux.HandleFunc("/crop/user", CheckState)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	server.ListenAndServe()
}

func CheckState(w http.ResponseWriter, r *http.Request) {
	fmt.Println("state 200")
}
