package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("vim-go")

	http.HandleFunc("/upload", UploadFile)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(33 << 22)

	var buf bytes.Buffer

	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")

	fmt.Println("Name: ", name)

	f, err := os.OpenFile("./downloaded", os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	io.Copy(f, file)

	contents := buf.String()
	fmt.Println(contents)
	buf.Reset()
}
