package rest

import (
	"commoles/db"
	"encoding/json"
	"net/http"
	"log"
)

func GetSteps(w http.ResponseWriter, r *http.Request) {
	SetCors(w, r)
	var con = db.CreateCon()
	result := db.GetAllSteps(con)
	log.Println("Get all logs")
	json.NewEncoder(w).Encode(result)
}
