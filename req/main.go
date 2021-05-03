package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
