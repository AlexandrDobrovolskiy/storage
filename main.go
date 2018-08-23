package main

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
	"storage/controllers"
	"log"
	"flag"
)

var Server *http.Server

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := mux.NewRouter()

	router.HandleFunc("/images/news", controllers.UploadImage).Methods("POST")
	router.HandleFunc("/files/news", controllers.UploadFile).Methods("POST")

	var dir string
	flag.StringVar(&dir, "dir", "public/", "../public/")
	flag.Parse()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir)))).Methods("GET")


	fmt.Println("Listening on ", port)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	Server = &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	log.Fatal(Server.ListenAndServeTLS(os.Getenv("CERT"), os.Getenv("KEY")))
}

