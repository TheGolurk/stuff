package main

import (
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

}
