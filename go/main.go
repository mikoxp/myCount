package main

import (
	rest "commoles/rest"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8030"
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Up and running...")
	})
	// logs
	router.HandleFunc("/logs", rest.GetLogs).Methods("GET")
	//steps
	router.HandleFunc("/steps", rest.GetSteps).Methods("GET")

	log.Printf("Serwer running in http://localhost:%s\n", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
