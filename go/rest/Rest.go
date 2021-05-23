package rest

import (
	"commoles/db"
	"encoding/json"
	"net/http"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
	var con = db.CreateCon()
	data := db.GetAll(con)
	n := data.Len()
	result := make([]db.Info, n)
	i := 0
	for temp := data.Front(); temp != nil; temp = temp.Next() {
		result[i] = temp.Value.(db.Info)
		i++
	}
	json.NewEncoder(w).Encode(result)
}
