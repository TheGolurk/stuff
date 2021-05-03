package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	user     = os.Getenv("tbot_user")
	password = os.Getenv("tbot_pwd")
	token    = ""
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

	// Login, get token and continue (only execute when auth is not working)
	response, err := http.Post("http://localhost:3001/agrosmart/v1/user/login", "application/json", bytes.NewBuffer(UserJson))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if response.StatusCode == http.StatusBadRequest {
		// return custom error
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	cookies := response.Cookies()
	for i, v := range cookies {
		fmt.Println(v, i)
		if v.Name == "Token" {
			token = v.String()
		}
	}

	// Then we can do whatever request that we need
}

func makeRequest(path, method string) error {
	url, err := url.Parse(path)
	if err != nil {
		return err
	}

	req := &http.Request{
		URL: url,
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	println(res)

	return nil
}
