package rest

import (
	"commoles/db"
	"encoding/json"
	"net/http"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
	SetCors(w, r)
	var con = db.CreateCon()
	result := db.GetAllLogs(con)
	json.NewEncoder(w).Encode(result)
}
