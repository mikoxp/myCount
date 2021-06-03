package rest

import (
	"commoles/db"
	"encoding/json"
	"net/http"
	"log"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
	SetCors(w, r)
	var con = db.CreateCon()
	result := db.GetAllLogs(con)
	log.Println("Get all logs")
	json.NewEncoder(w).Encode(result)
}
