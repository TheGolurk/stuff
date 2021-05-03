package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var (
	user     = os.Getenv("tbot_user")
	password = os.Getenv("tbot_pwd")
)

type UserInfo struct {
	User     string
	Password string
}

func main() {

	User := UserInfo{
		User:     user,
		Password: password,
	}

	UserJson, err := json.Marshal(&User)
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("/", "application/json", bytes.NewBuffer(UserJson))
	if err != nil {
		log.Fatal(err)
	}
}
