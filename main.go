package main

import (
	_ "api-repository-golang-v1/docs"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/image/bmp"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ImageStruct struct {
	imageName string `json:"imageName" example:"test.png"`
}

type User struct {
	userName string `json:"userName" example:"hello"`
}

// @title imagen repository
// @version 1.0
// @description Retriving from imagen repository
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/documentacion-digital/repo/imagenes-base64/{imageName}", getDocumentBase64).Methods("GET")
	router.HandleFunc("/documentacion-digital/repo/imagenes/{imageName}", getDocument).Methods("GET")

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// GetDocumentBase64 godoc
// @Summary Get image base64
// @Description Get image base64
// @Tags api-repo-image
// @Accept  json
// @Produce  json
// 200 {string} ImageStruct
// @Router /documentacion-digital/repo/imagenes-base64/{imageName} [get]
func getDocumentBase64(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	f, err := os.Open("/imagenes/" + params["imageName"])
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

// GetDocument godoc
// @Summary Get image
// @Description Get image
// @Tags api-repo-image
// @Accept  json
// @Produce  json
// 200 {string} ImageStruct
// @Router /documentacion-digital/repo/imagenes/{imageName} [get]
func getDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	reader, err := os.Open("/imagenes/" + params["imageName"])
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
		return
	}
	println("Got format String", formatString)

	if formatString == "png" {
		resizePNG := resize.Resize(100, 100, m, resize.Lanczos3)
		buff := new(bytes.Buffer)
		err = png.Encode(buff, resizePNG)
		if err != nil {
			println("failed to create buffer", err)
		}
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", strconv.Itoa(len(buff.Bytes())))
		if _, err := w.Write(buff.Bytes()); err != nil {
			log.Println("unable to write image.")
		}
	}
	if formatString == "jpeg" {
		resizeJPEG := resize.Resize(100, 100, m, resize.Lanczos3)
		buff := new(bytes.Buffer)
		err = jpeg.Encode(buff, resizeJPEG, nil)
		if err != nil {
			println("failed to create buffer", err)
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(buff.Bytes())))
		if _, err := w.Write(buff.Bytes()); err != nil {
			log.Println("unable to write image.")
		}
	}
	if formatString == "bmp" {
		resizeBMP := resize.Resize(100, 100, m, resize.Lanczos3)
		buff := new(bytes.Buffer)
		err = bmp.Encode(buff, resizeBMP)
		if err != nil {
			println("failed to create buffer", err)
		}
		w.Header().Set("Content-Type", "image/bmp")
		w.Header().Set("Content-Length", strconv.Itoa(len(buff.Bytes())))
		if _, err := w.Write(buff.Bytes()); err != nil {
			log.Println("unable to write image.")
		}
	}
}
