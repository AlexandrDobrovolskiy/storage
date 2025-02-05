package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, X-Access-Token, X-Application-Name, X-Request-Sent-Time")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
}

func RespondText(w http.ResponseWriter, data string) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, X-Access-Token, X-Application-Name, X-Request-Sent-Time, Accept-Encoding, X-Compress")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write([]byte(data))
}

func RespondFile(w http.ResponseWriter, r *http.Request, file []byte, filename string) {
	w.Header().Add("Content-Disposition", "inline")
	w.Header().Add("filename", filename)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, X-Access-Token, X-Application-Name, X-Request-Sent-Time")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	http.ServeContent(w, r, filename, time.Now(), bytes.NewReader(file))
}
