package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag/example/basic/docs"
	"log"
	"net/http"
)

type ImageStruct struct {
	//OrderID      string    `json:"orderId" example:"1"`
	//CustomerName string    `json:"customerName" example:"Leo Messi"`
	//OrderedAt    time.Time `json:"orderedAt" example:"2019-11-09T21:21:46+00:00"`
	//Items        []Item    `json:"items"`
	ImageName string `json:"imageName" example:"test.png"`
}

type User struct {
	userName string `json:"userName" example:"hello"`
}

// @title imagen repository
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()
	//router.HandleFunc("/documentacion-digital/repo/imagenes-base64/{name}", getDocumentBase64).Methods("GET")
	router.HandleFunc("/hello/{user}", getUser).Methods("GET")

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetUser godoc
// @Summary Get user name
// @Description Get user name
// @Tags api-repo-image
// @Accept  json
// @Produce  json
// 200 {string} User
// @Router /hello/{userName} [get]
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello")
}

/*
// GetDocumentBase64 godoc
// @Summary Get image base64
// @Description Get image base64
// @Tags api-repo-image
// @Accept  json
// @Produce  json
// 200 {object} ImageStruct
// @Router /documentacion-digital/repo/imagenes-base64/{name} [get]
func getDocumentBase64(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	f, err := os.Open("/imagenes/" + params["name"])
	if err != nil {
		log.Fatal("There are an error to open the file : ", err)
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal("Occurs and error to read the file : ", err)
	}
	encoded := base64.StdEncoding.EncodeToString(content)
	json.NewEncoder(w).Encode(encoded)
}

*/
