package rest

import (
	"commoles/db"
	"encoding/json"
	"log"
	"net/http"
)

func GetSteps(w http.ResponseWriter, r *http.Request) {
	var con = db.CreateCon()
	result := db.GetAllSteps(con)
	log.Println("Get all steps")
	json.NewEncoder(w).Encode(result)
}
