package rest

import (
	"commoles/db"
	"encoding/json"
	"net/http"
)

func GetSteps(w http.ResponseWriter, r *http.Request) {
	SetCors(w, r)
	var con = db.CreateCon()
	result := db.GetAllSteps(con)
	json.NewEncoder(w).Encode(result)
}
