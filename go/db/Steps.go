package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Step struct {
	DAY          string `json:"day"`
	STEPS_NUMBER int    `json:"stepsNumber"`
}

func GetAllSteps(con *sql.DB) []Step {
	values := make([]Step, 0)

	results, err := con.Query(GetAllStepsSelect)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var record Step
		err = results.Scan(&record.DAY, &record.STEPS_NUMBER)
		if err != nil {
			panic(err.Error())
		}
		values = append(values, record)
	}
	defer results.Close()

	return values
}
