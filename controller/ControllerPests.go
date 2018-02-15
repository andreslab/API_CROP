package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreslab/prj_api_crop/model"
)

func ManagerRouterPests(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		RequestGetPests(w, r)
	case "POST":
		fmt.Println("Resquest POST not wonrking")
	case "UPDATE":
		fmt.Println("Resquest UPDATE not wonrking")
	case "DELETE":
		fmt.Println("Resquest DELETE not wonrking")
		//RequestDeletePests(w, r)
	default:
	}
}

func RequestGetPests(w http.ResponseWriter, r *http.Request) {
	var data model.PestsModel

	/*db, err := database.NewMySQLDBPests()
	Pests, err := db.ListPests()
	if err != nil {
		fmt.Println("error")
	}
	for _, value := range Pests {
		data = append(data, *value)
	}*/

	data = model.PestsModel{
		ID:          0,
		Name:        "Viruela",
		Description: "Ataca: tallo, peciolo, pedúnculos florales, pero principalmente hojas, formando en la superficie de estas, manchas castañas con centro gris y con el borde teñido finamente de color amarillo. Si el ataque adquiere fuerza provoca defoliación, dejando los frutos expuestos al sol.",
		Treatment:   "Tratamiento e implementación de malla tutora para prevenir enfermedades fúngicas. Los fungicidas que se pueden emplear, dependiendo de la enfermedad fúngica a tratar son los siguientes: Azoxistrobina, Benadaxil, Beonomil, Boscalid, Bromuro de metilo, Captan, Carbendazil, Clorotalonil, Cyprodinil, Difenoconazole, Ferbam, Fludioxonil, Folpet, Flosetil aluminio, Hidróxido de cobre, Kasugamicina, Mancozeb, Metalaxil – m – isómero, Oxicloruro de cobre, Oxido cuproso, Procimidone, Propamocarb clorhidrato, Propineb, Pyraclostrobin, Sulfato cúprico pentahidratado, Sulfato tetracupico tricalcico, Tetraconazole, Tiram, Triadimefon, Zineb, Ziram.",
		Percent:     "10",
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
