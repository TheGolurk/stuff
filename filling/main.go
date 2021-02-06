package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("file.html")
	if err != nil {
		fmt.Println(err)
		panic("!!!!!!")
	}

	fmt.Println(string(data))
}
