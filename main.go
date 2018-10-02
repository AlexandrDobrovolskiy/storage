package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"FITstorage/config"
	"FITstorage/controllers"
	"runtime"
)

var Server *http.Server

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	envName := *flag.String("c", "server.cfg", "Environment config name")

	err := config.LoadConfig(envName)
	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()

	// Old upload
	router.HandleFunc("/images/news", controllers.UploadImage).Methods("POST")
	router.HandleFunc("/files/news", controllers.UploadFile).Methods("POST")

	// FilePond processing
	router.HandleFunc("/filepond", controllers.FilePondProcess).Methods("POST")
	router.HandleFunc("/filepond", controllers.FilePondDelete).Methods("DELETE")
	router.HandleFunc("/filepond", controllers.FilePondLoad).Methods("GET")

	// FilePond options
	router.HandleFunc("/filepond", controllers.Options).Methods("OPTIONS")

	// FilePond submit
	router.HandleFunc("/filepond/confirm", controllers.SubmitStore).Methods("POST")

	var dir string
	flag.StringVar(&dir, "dir", "public/", "public")
	flag.Parse()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir)))).Methods("GET")

	fmt.Println("Listening on ", config.Config.Server.Port)

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
		Addr:         config.Config.Server.Port,
		Handler:      router,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	if config.Config.Server.SecureConn {
		log.Fatal(Server.ListenAndServeTLS(config.Config.Server.SecureCert, config.Config.Server.SecureKey))
	} else {
		log.Fatal(Server.ListenAndServe())
	}
}
