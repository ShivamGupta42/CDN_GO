package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

type Error struct {
	ErrorReason string `json:"error_reason"`
}

func GetFiles(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)

	v, ok := m["fileName"]
	if !ok || v == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Error{"Null FileName"})
		return
	}

	f, err := os.ReadFile("./cachedFiles/" + strings.TrimSpace(v))

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(Error{"File Not Found"})
		//Call to underlying origin server in real CDN
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	w.Write(f)
}
