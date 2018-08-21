package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
	"storage/controllers"
	"log"
)

func main() {

	router := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router.HandleFunc("/images/news", controllers.UploadImage).Methods("POST")

	//test
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("YAY")
	}).Methods("GET")

	fmt.Println("Listening on ", port)

	log.Fatal(http.ListenAndServe("192.168.31.50:" + port, router))
}

