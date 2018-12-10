package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	recaptcha "github.com/dpapathanasiou/go-recaptcha"
)

/*Create new API Token*/
func CreateNewAPIToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[->] createToken")

	r.ParseForm()
	challenge := r.PostFormValue("g-recaptcha-response")
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	email := r.PostFormValue("email")

	if challenge == "" {
		w.WriteHeader(400)
		return
	}
	result, err := recaptcha.Confirm(ip, challenge)
	if err != nil {
		log.Println("recaptcha server error", err)
	}
	if result != true {
		w.WriteHeader(401)
		return
	}

	// Request is valied with google
	fmt.Println("	-> " + email)

}
