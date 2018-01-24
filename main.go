// TODO
// 1. Add spots data to MongoDB
// 2. Create API to return the Spot  given wind and tide direction

// TODO Some examples:
// https://thenewstack.io/make-a-restful-json-api-go/
// https://medium.com/@maumribeiro/a-fullstack-epic-part-i-a-rest-api-in-go-accessing-mongo-db-608b46e969cd
// https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3
// http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/
// http://himarsh.org/build-restful-api-microservice-with-go/
// http://www.blog.labouardy.com/build-restful-api-in-go-and-mongodb/

// Frameworks/Libraries
// gin-gonic
// gorm

// https://github.com/swrobel/meta-surf-forecast

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/meianuca/meianuca-go/models"
	"github.com/meianuca/meianuca-go/qualpraia"
)

// Models initialization
var regions = []models.Region{
	models.Region{ID: "florianopolis", Name: "Florianópolis"},
	models.Region{ID: "garopaba", Name: "Garopaba"},
	models.Region{ID: "imbituba", Name: "Imbituba"},
	models.Region{ID: "laguna", Name: "Laguna"},
}
var spots = []models.Spot{
	models.Spot{ID: "acores", Name: "Açores", RegionID: "florianopolis"},
	models.Spot{ID: "armacao", Name: "Armação", RegionID: "florianopolis"},
	models.Spot{ID: "barra-da-lagoa", Name: "Barra da Lagoa", RegionID: "florianopolis"},
	models.Spot{ID: "brava", Name: "Brava", RegionID: "florianopolis"},
	models.Spot{ID: "caldeira", Name: "Caldeirão", RegionID: "florianopolis"},
	models.Spot{ID: "campeche", Name: "Campeche", RegionID: "florianopolis"},
	models.Spot{ID: "galheta", Name: "Galheta", RegionID: "florianopolis"},
	models.Spot{ID: "galheta-norte", Name: "galheta-norte", RegionID: "florianopolis"},
	models.Spot{ID: "igrejinha", Name: "Igrejinha", RegionID: "florianopolis"},
	models.Spot{ID: "ingleses", Name: "Ingleses", RegionID: "florianopolis"},
	models.Spot{ID: "joaquina", Name: "Joaquina", RegionID: "florianopolis"},
	models.Spot{ID: "lagoinha-do-leste", Name: "lagoinha-do-leste", RegionID: "florianopolis"},
	models.Spot{ID: "lomba-do-sabao", Name: "Lomba do Sabão", RegionID: "florianopolis"},
	models.Spot{ID: "matadeiro", Name: "Matadeiro", RegionID: "florianopolis"},
	models.Spot{ID: "mocambique", Name: "Moçambique", RegionID: "florianopolis"},
	models.Spot{ID: "mole", Name: "Mole", RegionID: "florianopolis"},
	models.Spot{ID: "morro-das-pedras", Name: "Morro das Pedras", RegionID: "florianopolis"},
	models.Spot{ID: "naufragados", Name: "Naufragados", RegionID: "florianopolis"},
	models.Spot{ID: "novo-campeche", Name: "Novo Campeche", RegionID: "florianopolis"},
	models.Spot{ID: "pico-da-cruz", Name: "Pico da Cruz", RegionID: "florianopolis"},
	models.Spot{ID: "riozinho", Name: "Riozinho", RegionID: "florianopolis"},
	models.Spot{ID: "santinho", Name: "Santinho", RegionID: "florianopolis"},
	models.Spot{ID: "solidao", Name: "Solidão", RegionID: "florianopolis"},
}

// APIs

// GetAllSpots gets all spots
func GetAllSpots(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	// log.Printf("Query: wind=%s / swell=%s", queries["wind"], queries["swell"])

	wind, swell := queries["wind"][0], queries["swell"][0]
	var spots = qualpraia.FindByWindAndSwell(wind, swell)

	json.NewEncoder(w).Encode(spots)
}

// GetSpot gets a single Spot given its ID
func GetSpot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range spots {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// GetAllRegions return all the regions
func GetAllRegions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(regions)
}

// Main function
func main() {
	router := mux.NewRouter() // Router initialization

	// Spots
	// router.HandleFunc("/spots", GetAllSpots).Methods("GET")
	router.HandleFunc("/spots", GetAllSpots).Methods("GET").Queries("wind", "{wind:[a-zA-Z]{1,2}}", "swell", "{swell:[a-zA-Z]{1,2}}")
	router.HandleFunc("/spots/{id}", GetSpot).Methods("GET")
	// Regions
	router.HandleFunc("/regions", GetAllRegions).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
