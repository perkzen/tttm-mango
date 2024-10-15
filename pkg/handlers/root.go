package handlers

import "net/http"

func HandleRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "This is tttm-tango by Domen Perko"}`))
}
