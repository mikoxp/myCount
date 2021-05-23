package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Info struct {
	ID   int    `json:"id"`
	INFO string `json:"info"`
}

func GetAll(con *sql.DB) []Info {
	values := make([]Info, 0)

	results, err := con.Query(GetAllLogs)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var info Info
		err = results.Scan(&info.ID, &info.INFO)
		if err != nil {
			panic(err.Error())
		}
		values = append(values, info)
	}
	defer results.Close()

	return values
}
