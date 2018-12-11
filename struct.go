package main

type APIToken struct {
	UID     int    `json:"uid"`
	Email   string `json:"email"`
	Token   string `json:"token"`
	blocked int    `json:"blocked"`
}
