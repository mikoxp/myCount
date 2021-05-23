package main

import (
	rest "commoles/rest"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Serwer running in http://localhost:8083")

	http.HandleFunc("/logs", rest.GetLogs)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
