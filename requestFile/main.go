package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("vim-go")

	http.HandleFunc("/upload", UploadFile)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

}
