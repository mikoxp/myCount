package rest

import (
	"commoles/db"
	"encoding/json"
	"log"
	"net/http"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
	var con = db.CreateCon()
	result := db.GetAllLogs(con)
	log.Println("Get all logs")
	json.NewEncoder(w).Encode(result)
}
