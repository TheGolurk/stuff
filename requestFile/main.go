package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("vim-go")

	http.HandleFunc("/upload", UploadFile)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(33 << 22)

	buffer := bytes.Buffer

	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	name := strings.Split(header.Filename, ".")

	fmt.Println("Name: ", name)

}
