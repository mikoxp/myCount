package rest

import (
	"commoles/db"
	"encoding/json"
	"net/http"
)

func GetSteps(w http.ResponseWriter, r *http.Request) {
	var con = db.CreateCon()
	result := db.GetAllSteps(con)
	json.NewEncoder(w).Encode(result)
}