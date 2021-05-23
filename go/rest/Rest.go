package rest

import (
	"commoles/db"
	"encoding/json"
	"net/http"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
	var con = db.CreateCon()
	result := db.GetAll(con)
	json.NewEncoder(w).Encode(result)
}
