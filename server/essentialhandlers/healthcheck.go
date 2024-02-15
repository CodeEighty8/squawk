package essentialhandlers

import (
	"net/http"
	"time"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(time.Now().UTC().String()))
}
