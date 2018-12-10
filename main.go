package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	recaptcha "github.com/dpapathanasiou/go-recaptcha"
	"github.com/gorilla/mux"
)

/*
	Variablen:

	PORT
	recaptchaPrivateKey
	aws_id
	aws_secert
*/

func main() {
	fmt.Println("[Start] API on Port :: ", os.Getenv("PORT"))

	//Google recaptcha
	recaptcha.Init(os.Getenv("recaptchaPrivateKey"))

	//Webserver
	router := mux.NewRouter()
	router.HandleFunc("/api/token", CreateNewAPIToken).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
