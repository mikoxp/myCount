package endpoint

import "net/http"

func SetCors(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
