package main

import (
	"fmt"
	"net/http"
)

/*Create new API Token*/
func CreateNewAPIToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[->] createToken")

	r.ParseForm()
	challenge := r.PostFormValue("g-recaptcha-response")
	//ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	email := r.PostFormValue("email")
	if challenge == "" {
		w.WriteHeader(400)
		return
	}
	/*
		result, err := recaptcha.Confirm(ip, challenge)
		if err != nil {
			log.Println("recaptcha server error", err)
		}
		if result != true {
			w.WriteHeader(401)
			return
		}
	*/
	// Request is valied with google
	if email == "" {
		w.WriteHeader(400)
		return
	}
	fmt.Println("	-> " + email)

	cmail := checkIfEmailExist(email)
	fmt.Println("	-> Email exist status :: ", cmail)
	if cmail {
		w.WriteHeader(409)
		return
	}

	token := generateToken(email)
	fmt.Println("	-> new token :: " + token)
	createNewAPIToken(email, token)
	sendSubmittMessage(email, token)

}
